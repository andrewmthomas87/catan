FROM golang:1.15-alpine3.13 as builder

WORKDIR /go/src/github.com/andrewmthomas87/catan

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -v ./cmd/auto_migrate

FROM alpine:3.13

# wait-for-it requires bash, alpine doesn't ship with bash, so install it
RUN apk add --no-cache bash

RUN mkdir -p /go/src/github.com/andrewmthomas87/catan
WORKDIR /go/src/github.com/andrewmthomas87/catan

COPY --from=builder /go/src/github.com/andrewmthomas87/catan/auto_migrate .
COPY --from=builder /go/src/github.com/andrewmthomas87/catan/wait-for-it.sh .

RUN chmod +x wait-for-it.sh

CMD ["/go/src/github.com/andrewmthomas87/catan/auto_migrate"]
