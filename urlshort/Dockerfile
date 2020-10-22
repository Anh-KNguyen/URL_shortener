FROM golang:latest

WORKDIR /src/
COPY . /src/
RUN CGO_ENABLED=0 go build -o /bin/tinyurl

FROM scratch
COPY --from=build /bin/tinyurl /bin/tinyurl
ENTRYPOINT ["/bin/tinyurl"]