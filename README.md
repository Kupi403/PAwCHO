# Dokumentacja wdrożenia łańcucha CI/CD (Zadanie 2)
**Autor:** Michał Kupidura

## Opis zrealizowanego łańcucha CI (GitHub Actions)
Potok CI został zdefiniowany w pliku `.github/workflows/gha_example.yml` i realizuje następujące etapy:

1. Pobranie kodu źródłowego z repozytorium przy użyciu akcji `actions/checkout`.
2. Skonfigurowanie środowiska emulacji QEMU oraz silnika Buildx do obsługi kompilacji wieloarchitekturowej.
3. Logowanie do rejestru Docker Hub za pomocą tajnego tokenu PAT (`secrets.DOCKERHUB_TOKEN`) oraz nazwy użytkownika (`vars.DOCKERHUB_USERNAME`).
4. Logowanie do rejestru GitHub Container Registry (`ghcr.io`) przy użyciu automatycznego tokenu uwierzytelniającego `secrets.GITHUB_TOKEN`.
5. Wygenerowanie metadanych i automatycznych tagów OCI na podstawie hashu zatwierdzenia.
6. Zbudowanie obrazu lokalnie w celach weryfikacyjnych (bez wysyłania do rejestru).
7. Wykonanie testu CVE przy użyciu skanera Trivy. Wykrycie podatności o statusie CRITICAL lub HIGH przerywa działanie potoku z kodem wyjścia 1, blokując publikację obrazu.
8. Finalne zbudowanie obrazu dla architektur `linux/amd64` oraz `linux/arm64`, pobranie i przesłanie danych cache w trybie `max` do Docker Hub oraz wypchnięcie gotowego obrazu na `ghcr.io`.

## Przyjęty schemat tagowania obrazów i danych cache

1. **Obraz produkcyjny (ghcr.io):** Tagowany unikalnym, skróconym hashem commita (`sha-`). Zastosowanie niezmiennych identyfikatorów (paradygmat Immutable Infrastructure) zamiast zmiennego tagu `latest` gwarantuje determinizm wdrożeń i umożliwia jednoznaczne powiązanie obrazu z konkretną rewizją kodu.
2. **Dane pamięci podręcznej (Docker Hub):** Tagowane stałą wartością `:latest`. W bezstanowych środowiskach uruchomieniowych (ephemeral runners) stały tag ułatwia silnikowi BuildKit natychmiastowe pobranie ostatniego stanu cache bez konieczności dynamicznego mapowania struktur z poprzednich wywołań.
3. **Tryb zapisu cache (mode=max):** Z uwagi na architekturę wieloetapową (Multi-stage build), domyślny tryb `min` zapisałby wyłącznie warstwy końcowe (etap `FROM scratch`). Tryb `max` wymusza buforowanie kroków pośrednich (w tym kompilacji Go i optymalizacji UPX w etapie budującym), co pozwala skrócić czas kolejnych procesów budowy z kilkunastu minut do kilkudziesięciu sekund.

## Linki do zasobów
* Repozytorium GitHub: https://github.com/Kupi403/PAwCHO/tree/zadanie-2
* Rejestr obrazów (GHCR): https://ghcr.io/kupi403/weather-app
* Rejestr danych cache (Docker Hub): https://hub.docker.com/r/mickupi77/weather-app-cache