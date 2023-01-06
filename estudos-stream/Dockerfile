FROM golang:latest
WORKDIR /app
COPY . .
RUN go get gocdn
EXPOSE 8000
CMD ["go", "run", "main.go"]
