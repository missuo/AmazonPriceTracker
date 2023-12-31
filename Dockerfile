FROM golang:1.21 AS builder

WORKDIR /go/src/github.com/missuo/AmazonPriceTracker
COPY main.go ./
COPY go.mod ./
COPY go.sum ./
RUN go get -d -v ./
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o amazonpricetracker .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /go/src/github.com/missuo/AmazonPriceTracker/amazonpricetracker /app/amazonpricetracker
CMD ["/app/amazonpricetracker"]
