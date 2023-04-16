FROM golang:1.20.3-bullseye AS build
WORKDIR /usr/project/vpn
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY .env /usr/project/vpn

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /vpn
FROM alpine:3.17.3
USER root
RUN apk update && \
    apk add openvpn
RUN touch config.ovpn
COPY --from=build /vpn /usr/project/vpn
ENV PATH="${PATH}:/usr/sbin"
EXPOSE 8000
CMD ["/usr/project/vpn"]
