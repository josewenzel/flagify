FROM golang:1.20

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -v -o /flagify-app

EXPOSE 8123

CMD ["/flagify-app"]
