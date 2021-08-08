FROM golang:latest as BUILD
RUN apt-get clean && \
    apt-get update && \
    apt-get install -y binutils
WORKDIR builddir
COPY . .
ENV DBHOST=localhost
ENV DBPORT=5434
ENV DBUSER=postgres
ENV DBNAME=go-mux-db
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o webcache-tutorial

FROM scratch
COPY --from=BUILD ./go/builddir/webcache-tutorial .
ENTRYPOINT ["./webcache-tutorial"]
