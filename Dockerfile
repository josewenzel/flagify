FROM golang:1.20

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -C src -o /flagify-app

EXPOSE 8123

CMD ["/flagify-app"]
