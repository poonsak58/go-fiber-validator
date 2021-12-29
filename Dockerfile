FROM golang:alpine as build
WORKDIR /app
COPY main.go .
COPY go.mod .
COPY go.sum .
COPY . .
RUN go mod download
RUN go build -o api-server .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build /app/api-server .
EXPOSE 2000
CMD ["/app/api-server"]
