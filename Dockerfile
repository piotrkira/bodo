FROM golang:1.23-alpine AS build

ARG TARGETARCH

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOARCH=${TARGETARCH} go build -o bodo .

FROM alpine:latest
LABEL org.opencontainers.image.source=https://github.com/piotrkira/bodo
LABEL org.opencontainers.image.description="My container image"
LABEL org.opencontainers.image.licenses=MIT

WORKDIR /app
RUN mkdir -p /usr/local/share/bodo
RUN mkdir -p /etc/bodo
EXPOSE 8080
COPY config.yaml /etc/bodo
COPY index.html themes.yaml /usr/local/share/bodo

COPY --from=build /app/bodo .

ENTRYPOINT ["/app/bodo"]
