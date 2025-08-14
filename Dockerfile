# build stage
FROM swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/library/golang:1.23.4-alpine3.20 AS builder
ARG GOPROXY=https://goproxy.cn,direct
WORKDIR /src
COPY . .
RUN go build -ldflags '-s -w'

# server image

FROM swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/library/alpine:3.20
LABEL org.opencontainers.image.source https://github.com/go-shiori/shiori
COPY --from=builder /src/shiori /usr/bin/
RUN apk add --no-cache ca-certificates tzdata
USER root
WORKDIR /shiori
EXPOSE 8080
ENV SHIORI_DIR /shiori/
ENTRYPOINT ["/usr/bin/shiori"]
VOLUME [ "/shiori" ]
CMD ["server"]