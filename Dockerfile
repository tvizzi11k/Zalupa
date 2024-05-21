FROM golang:alpine AS BUILD
RUN mkdir /app
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o behappy

FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=BUILD /app/behappy .
CMD ["./behappy"]
