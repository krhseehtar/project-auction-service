# Build Stage

FROM golang:1.20.2 AS build

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o supply-side-service .

CMD ["go","run","main.go"]
# Final Stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=build /app/supply-side-service .

EXPOSE 8080

CMD ["./supply-side-service"]
