docker build -t forum .
docker container run -p 8080:8080 --detach --name dx forum