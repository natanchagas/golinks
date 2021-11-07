docker netwrok create golinks

docker build -t natanchagas/golinks-db:alpha database/.
docker run -it --name golinks-db --network golinks --rm -p 5432:5432 natanchagas/golinks-db:alpha

docker build -t natanchagas/golinks-api api/.
docker run -it --name golinks-api --network golinks --rm -p 8081:1323 -e DB_HOSTNAME="golinks-db" -e DB_USERNAME="golinksadm" -e DB_PASSWORD="mdasknilog" -e DB_TABLE="golinks" natanchagas/golinks-api:latest

docker build -t natanchagas/golinks-front front/.
docker 