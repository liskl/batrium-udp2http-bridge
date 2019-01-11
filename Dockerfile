# build stage
FROM golang:1.11.4-alpine AS build-env

RUN apk update && apk upgrade && \
    apk add --no-cache bash git

ARG RELEASE=unknown
ARG COMMIT=unknown
ARG BUILD_TIME=unknown

ENV PROJECT "gogs.infra.liskl.com/liskl/batrium-udp-listener"
ENV RELEASE ${RELEASE}
ENV COMMIT ${COMMIT}
ENV BUILD_TIME ${BUILD_TIME}

COPY ./src /go/src

WORKDIR /go/src/gogs.infra.liskl.com/liskl/batrium-udp-listener

RUN cd ./UDPmodule \
    && go install ; \
    cd .. \
    && go get \
    && go build \
  		-ldflags "-s -w \
      -X ${PROJECT}/main.Release=${RELEASE} \
  		-X ${PROJECT}/main.Commit=${COMMIT} \
      -X ${PROJECT}/main.BuildTime=${BUILD_TIME}" \
  		-o batrium-udp-listener ;


# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/gogs.infra.liskl.com/liskl/batrium-udp-listener/batrium-udp-listener /app/
CMD ["/app/batrium-udp-listener"]
