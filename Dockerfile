FROM golang:latest

WORKDIR /src/
COPY . /src/

CMD ["go", "run", "src/main.go"]