FROM golang:alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on

COPY . .

RUN cd shippy-cli-consignment && go mod download
RUN cd shippy-cli-consignment && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shippy-cli-consignment

# Run container
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/shippy-cli-consignment/shippy-cli-consignment .
COPY --from=builder /app/shippy-cli-consignment/consignment.json .

CMD ["./shippy-cli-consignment"]