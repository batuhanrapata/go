# PostgreSQL'in resmi Docker imajını kullan
FROM postgres:latest

# PostgreSQL veritabanı için gerekli ortam değişkenlerini ayarla
ENV POSTGRES_USER=myuser
ENV POSTGRES_PASSWORD=depixen-pass
ENV POSTGRES_DB=postgres

# PostgreSQL'in dinleyeceği portu aç
EXPOSE 5439

# PostgreSQL'i başlat
CMD ["postgres"]
