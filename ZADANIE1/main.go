package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"
)

const (
	AuthorName = "Michał Kupidura"
	Port       = "8080"
)

func main() {
	isCheck := flag.Bool("check", false, "Healthcheck mode")
	flag.Parse()

	if *isCheck {
		resp, err := http.Get("http://localhost:" + Port + "/health")
		if err != nil || resp.StatusCode != 200 { os.Exit(1) }
		os.Exit(0)
	}

	fmt.Printf("START: [%s] Autor: %s, Port: %s\n", time.Now().Format(time.RFC3339), AuthorName, Port)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(200) })

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var weatherData string
		if r.Method == http.MethodPost {
			city := r.FormValue("city")
			coords := map[string]string{
				"Lublin": "latitude=51.25&longitude=22.57",
				"Warszawa": "latitude=52.23&longitude=21.01",
			}[city]

			apiURL := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?%s&current_weather=true", coords)
			resp, err := http.Get(apiURL)
			if err == nil {
				defer resp.Body.Close()
				var data map[string]interface{}
				json.NewDecoder(resp.Body).Decode(&data)
				curr := data["current_weather"].(map[string]interface{})
				weatherData = fmt.Sprintf("Miasto: %s | Temp: %v°C | Wiatr: %v km/h", city, curr["temperature"], curr["windspeed"])
			} else {
				weatherData = "Błąd pobierania danych."
			}
		}

		tmpl := template.Must(template.New("web").Parse(`
			<!DOCTYPE html>
			<html lang="pl">
			<head>
				<meta charset="UTF-8"><meta name="viewport" content="width=device-width, initial-scale=1.0">
				<style>
					body { background: #0f172a; color: #f8fafc; font-family: system-ui; display: flex; justify-content: center; align-items: center; min-height: 100vh; margin: 0; }
					.card { background: #1e293b; padding: 2.5rem; border-radius: 1.5rem; box-shadow: 0 25px 50px -12px rgba(0,0,0,0.5); text-align: center; width: 320px; border: 1px solid #334155; }
					h1 { color: #38bdf8; font-size: 1.8rem; margin-bottom: 1.5rem; }
					select, button { width: 100%; padding: 0.75rem; border-radius: 0.5rem; border: none; margin-bottom: 1rem; font-size: 1rem; }
					select { background: #334155; color: white; }
					button { background: #0ea5e9; color: white; font-weight: 700; cursor: pointer; transition: 0.3s; }
					button:hover { background: #0284c7; transform: scale(1.02); }
					.res { background: #0c4a6e; padding: 1rem; border-radius: 0.5rem; border: 1px solid #0ea5e9; margin-top: 1rem; font-size: 0.9rem; }
					.footer { margin-top: 2rem; font-size: 0.75rem; color: #64748b; }
				</style>
				<title>Weather v3</title>
			</head>
			<body>
				<div class="card">
					<h1>Pogoda</h1>
					<form method="POST">
						<select name="city"><option>Lublin</option><option>Warszawa</option></select>
						<button type="submit">Sprawdź</button>
					</form>
					{{if .Result}}<div class="res">{{.Result}}</div>{{end}}
					<div class="footer">Student: {{.Author}}</div>
				</div>
			</body>
			</html>
		`))
		tmpl.Execute(w, struct{ Author, Result string }{AuthorName, weatherData})
	})

	http.ListenAndServe(":"+Port, nil)
}