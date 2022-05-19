FROM --platform=linux/amd64 golang:1.18-alpine as builder
WORKDIR /app
ADD . /app
RUN go mod vendor 
RUN go build -buildvcs=false -o /app-linux

FROM --platform=linux/amd64 alpine:latest
COPY --from=builder /app-linux /root/app
RUN chmod +x /root/app
RUN apk add ca-certificates

WORKDIR /root

ENTRYPOINT [ "./app" ]

ADD botbg.png /root/botbg.png
ADD assets /root/assets