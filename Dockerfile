# Obraz ubuntu najnowszy
FROM ubuntu:latest
# Dane autora
LABEL mainteiner="Michał Kupidura <s101607@pollub.edu.pl>"
# Aktualizacja systemu i instalacja Apache w jednej warstwie
RUN apt-get update && apt-get upgrade -y && apt-get install -y apache2 && apt-get clean
# kopioweanie pliku 
COPY index.html /var/www/html/index.html
#na porcie 80
EXPOSE 80
# Uruchomienie Apache w tle, aby kontener nie zgasł zaraz po starcie
CMD ["apache2ctl", "-D", "FOREGROUND"]