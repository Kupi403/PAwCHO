# syntax=docker/dockerfile:1
# ETAP 1: Builder
FROM golang:1.21-alpine AS builder

# Instalacja UPX i certyfikatów
RUN apk add --no-cache ca-certificates upx

WORKDIR /src
# Optymalizacja cache: kopiujemy tylko pliki definicji (jeśli masz go.mod)
# COPY go.mod go.sum ./
# RUN go mod download

COPY main.go .

# -s -w: usuwa debug i tablice symboli
# -trimpath: usuwa ścieżki z systemu buildera z binarki (bezpieczeństwo + rozmiar)
# -buildvcs=false: nie dodaje informacji o Gicie do binarki
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-s -w -extldflags '-static'" \
    -trimpath \
    -buildvcs=false \
    -o weather-app main.go && \
    upx --ultra-brute weather-app

# Tworzymy minimalny plik passwd (tylko 1 linia zamiast całego systemu)
RUN echo "app:x:10001:10001::/:" > /etc/passwd.min

# ETAP 2: Finalny obraz
FROM scratch

# Metadane OCI
LABEL org.opencontainers.image.authors="Michał Kupidura" \
      org.opencontainers.image.title="WeatherApp-Pro-Minimal"

# Kopiujemy minimum
COPY --from=builder /etc/passwd.min /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /src/weather-app /weather-app

USER 10001
EXPOSE 8080

# Healthcheck zostawiamy , ale w scratch i tak waży 0B w warstwach
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD ["/weather-app", "-check"]

ENTRYPOINT ["/weather-app"]