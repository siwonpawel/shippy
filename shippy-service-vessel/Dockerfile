FROM golang:alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on

COPY . .

RUN cd shippy-service-vessel && go mod download
RUN cd shippy-service-vessel && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shippy-service-vessel

# Run container
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/shippy-service-vessel/shippy-service-vessel .

CMD ["./shippy-service-vessel"]