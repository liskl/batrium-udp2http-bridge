# build stage
FROM golang:1.11.4-alpine AS build-env

RUN apk update && apk upgrade && \
    apk add --no-cache bash git

ARG RELEASE=unknown
ARG COMMIT=unknown
ARG BUILD_TIME=unknown

ENV PROJECT "github.com/liskl/batrium-udp2http-bridge"
ENV RELEASE ${RELEASE}
ENV COMMIT ${COMMIT}
ENV BUILD_TIME ${BUILD_TIME}

COPY ./main.go /go/src/github.com/liskl/batrium-udp2http-bridge/main.go
COPY ./batrium /go/src/github.com/liskl/batrium-udp2http-bridge/
COPY ./UDPmodule /go/src/github.com/liskl/batrium-udp2http-bridge/

WORKDIR /go/src/github.com/liskl/batrium-udp2http-bridge

RUN cd /go/src/github.com/liskl/batrium-udp2http-bridge/UDPmodule \
    && go install ; \
    cd .. \
    && go get github.com/liskl/batrium-udp2http-bridge \
    && go build \
  		-ldflags "-s -w \
      -X ${PROJECT}/main.Release=${RELEASE} \
  		-X ${PROJECT}/main.Commit=${COMMIT} \
      -X ${PROJECT}/main.BuildTime=${BUILD_TIME}" \
  		-o batrium-udp2http-bridge ;

RUN find /go/src -type d

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/github.com/liskl/batrium-udp2http-bridge/batrium-udp2http-bridge /app/
CMD ["/app/batrium-udp2http-bridge"]
