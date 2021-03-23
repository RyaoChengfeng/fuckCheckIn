FROM golang:latest AS builder
COPY ./src /src
WORKDIR /src
ENV GOPROXY=https://goproxy.io GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o /build/app

FROM alpine:latest
COPY ./env /env
COPY --from=builder /build/app .
ENTRYPOINT ["sh","-c","/app"]
