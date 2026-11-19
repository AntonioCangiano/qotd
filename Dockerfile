# DA USARE SOLO SE LO SCAN TRIVY BLOCCA LA PIPELINE.
# Multi-stage: si compila nell'immagine golang (grande, con molti pacchetti)
# ma si spedisce solo il binario su una base minima Red Hat.
FROM docker.io/library/golang:1.23 AS build
WORKDIR /app
COPY *.go ./
RUN go mod init qotd && go get github.com/gorilla/mux github.com/gorilla/handlers \
 && CGO_ENABLED=0 go build -o qotd .

FROM registry.access.redhat.com/ubi9/ubi-micro:latest
COPY --from=build /app/qotd /qotd
EXPOSE 10000
USER 1001
CMD ["/qotd"]
