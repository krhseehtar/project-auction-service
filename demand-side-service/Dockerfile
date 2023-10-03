# Build Stage
FROM golang:1.20.2 AS build

ENV MYSQL_HOST=${MYSQL_HOST}

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o demand-side-service .

CMD ["go","run","main.go"]

# Final Stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=build /app/demand-side-service .

EXPOSE 8081

CMD ["./demand-side-service"]
