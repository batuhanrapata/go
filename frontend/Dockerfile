# Adım 1: Geliştirme aşaması
FROM node:20

# Çalışma dizinini ayarla
WORKDIR /app

# Bağımlılıkları yükle
COPY package*.json ./
RUN npm install

# Uygulama kaynak kodunu kopyala
COPY . .

# Geliştirme ortamında çalışacak uygulama için port 3000'i aç
EXPOSE 3000

# Uygulamayı başlat
CMD ["npm", "start"]
