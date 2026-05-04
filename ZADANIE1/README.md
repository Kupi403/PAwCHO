# Aplikacja Pogodowa - Zadanie 1 (PAwCHO)

Projekt realizowany w ramach laboratorium "Programowanie Aplikacji w Chmurze Obliczeniowej" z wykorzystaniem języka **Go**,.

### Najważniejsze cechy:

- **Rozmiar obrazu:** ~2.4 MB (dzięki kompresji UPX i bazie `scratch`).
- **Bezpieczeństwo:** 0 podatności w Docker Scout, praca w trybie `non-root`.
- **Multi-arch:** Wsparcie dla `amd64` oraz `arm64`.
- **Interfejs:** Nowoczesny "Dark Mode" z obsługą API Open-Meteo.

## 📄 Pełna Dokumentacja

Szczegółowy opis realizacji zadania, analiza bezpieczeństwa oraz zrzuty ekranu znajdują się w pliku:
👉 [**opis.pdf**](./opis.pdf)

## 🛠 Jak uruchomić lokalnie?

```bash
# Pobranie obrazu z Docker Hub
docker pull mickupi77/weather-app:v5

# Uruchomienie kontenera
docker run -d --name weather-app -p 8081:8080 mickupi77/weather-app:v5
```
