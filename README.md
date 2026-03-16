Generowanie certyfikatów bezpieczeństwa: Przy użyciu narzędzia OpenSSL wygenerowano klucz prywatny RSA (4096-bit) oraz samopodpisany certyfikat X.509. Był to niezbędny krok do zabezpieczenia komunikacji z lokalnym rejestrem obrazów poprzez protokół HTTPS (TLS).

Konfiguracja i uruchomienie lokalnego rejestru: Wykorzystano oficjalny obraz registry:2. Poprzez mechanizm wolumenów (Bind Mounts) przekazano certyfikaty do kontenera, a za pomocą zmiennych środowiskowych wymuszono na serwerze pracę w bezpiecznym trybie TLS na porcie 443.

Optymalizacja Docker Engine: W konfiguracji Docker Desktop dodano adres 127.0.0.1:443 do listy insecure-registries. Pozwoliło to na autoryzację połączenia z rejestrem mimo użycia certyfikatu self-signed (niepodpisanego przez publiczne centrum certyfikacji).

Opracowanie zoptymalizowanego pliku Dockerfile: Stworzono instrukcję budowania obrazu web100 opartą na Ubuntu. Zastosowano dobre praktyki, takie jak łączenie komend instalacji (apt-get) w jedną warstwę i czyszczenie plików tymczasowych, co zredukowało objętość obrazu i liczbę warstw w systemie plików UnionFS.

Budowa i tagowanie obrazu: Zbudowano obraz web100 zawierający serwer Apache oraz stronę z danymi studenta. Następnie obraz został oznaczony tagiem wskazującym na adres lokalnego rejestru, co umożliwiło jego poprawną identyfikację przez klienta Docker.

Publikacja obrazu w rejestrze (Push): Wykonano operację przesłania obrazu do lokalnego serwera. Proces ten zweryfikował poprawność działania bezpiecznego połączenia i poprawną strukturę warstw (blobs) w magazynie rejestru.
