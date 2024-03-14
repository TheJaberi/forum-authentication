FROM golang:1.21.4
LABEL authors="A4H"
LABEL version="0.1"
COPY . /forum
WORKDIR /forum/main
RUN go build -v main.go
EXPOSE 8080
CMD ./main