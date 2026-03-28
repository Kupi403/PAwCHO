# build
FROM alpine:3.19 AS builder

ARG VERSION=1.0
ENV APP_VERSION=$VERSION

WORKDIR /app
COPY index.sh .
RUN chmod +x index.sh  # nadajemy prawa już tutaj

# scratch
FROM scratch AS stage1
# tylko kopiujemy plik, nie robimy RUN
COPY --from=builder /app/index.sh /index.sh

# nginx
FROM nginx:alpine

ARG VERSION
ENV APP_VERSION=$VERSION

RUN apk add --no-cache curl bash

WORKDIR /usr/share/nginx/html

# kopiujemy skrypt z scratch
COPY --from=stage1 /index.sh ./index.sh

EXPOSE 80

# uruchamiamy skrypt dynamicznie przy starcie kontenera
CMD ["sh", "-c", "./index.sh && nginx -g 'daemon off;'"]

HEALTHCHECK --interval=30s --timeout=5s --start-period=5s \
  CMD curl -f http://localhost || exit 1