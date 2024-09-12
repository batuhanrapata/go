backend projesini böyle çalıştırılmalı.

```bash
docker build -t go-backend .
docker run --name backend-container --link postgres-db:postgres-db -e DB_HOST=postgres-db -e DB_PORT=5439 -e DB_USER=postgres -e DB_PASSWORD=depixen-pass -e DB_NAME=postgres -p 8080:8080 go-backend
```

frontend projesi böyle çalıştırılmalı

```bash
docker build -t my-react-app .
docker run -p 80:80 my-react-app
```
