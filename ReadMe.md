docker build -t my-postgres .  
docker run -d -p 5439:5432 --name my-postgres-container -v depixen-volume:/var/lib/postgresql/data my-postgres

docker build -t my-go-app .  
docker run -d --name my-go-app-container -p 8080:8080 --link my-postgres-container my-go-app

docker build -t my-react-app .  
docker run -d -p 3000:3000 my-react-app
