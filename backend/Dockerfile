# 1. Aşama: Go Uygulamasını Derleme
FROM golang:1.23-alpine AS go-builder

# Çalışma dizini ayarla
WORKDIR /app

# Go mod dosyalarını kopyala ve bağımlılıkları indir
COPY go.mod go.sum ./
RUN go mod download

# Uygulama kaynak kodunu kopyala
COPY . .

# Uygulamayı derle
RUN go build -o main ./cmd/api

# 2. Aşama: Minimal Image Üzerinde Çalıştırma
FROM alpine:latest

# Çalışma dizini ayarla
WORKDIR /root/

# Derlenmiş uygulamayı minimal image'a kopyala
COPY --from=go-builder /app/main .

VOLUME /root/uploads /root/uploads

# PostgreSQL için gerekli ortam değişkenlerini ayarla
ENV DB_HOST=postgres-db
ENV DB_PORT=5439
ENV DB_USER=postgres
ENV DB_PASSWORD=123qwe
ENV DB_NAME=postgres

# Gerekli dosyayı kopyala
COPY firestorage-45220-firebase-adminsdk-eiqhg-1f718a9e75.json /root/
COPY .env /root/

# Uygulamayı dinleyecek portu aç
EXPOSE 8080

# Uygulamayı başlat
CMD ["./main"]
