# syntax=docker/dockerfile:1
FROM golang:1.16

WORKDIR /src
ENV CGO_ENABLED=0
COPY go.* .
RUN go mod download
COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build go build -a -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /src/app ./
CMD ["./app"]
EXPOSE 8080
