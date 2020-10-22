FROM golang:latest

COPY . /usr/share/nginx/html

CMD ["go", "run", "src/main.go"]