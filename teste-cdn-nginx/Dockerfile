FROM golang:latest

COPY main.go .

RUN go build -o app main.go

RUN apt-get update && apt-get install -y nginx

COPY nginx.conf /etc/nginx/nginx.conf

CMD ["nginx", "-g", "daemon off;"]
CMD ["./app"]
