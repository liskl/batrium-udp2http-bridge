# build stage
FROM golang:1.16-alpine AS build-env

RUN apk update && apk upgrade && \
    apk add --no-cache bash git

ENV PROJECT "github.com/liskl/batrium-udp2http-bridge"

COPY ./main.go /go/src/${PROJECT}/main.go
COPY ./go.sum /go/src/${PROJECT}/go.sum
COPY ./go.mod /go/src/${PROJECT}/go.mod

COPY ./static /go/src/${PROJECT}/static
COPY ./templates /go/src/${PROJECT}/templates
#COPY ./vendor /go/src/${PROJECT}/vendor


COPY ./batrium /go/src/${PROJECT}/batrium
COPY ./metrics /go/src/${PROJECT}/metrics


WORKDIR /go/src/${PROJECT}

ARG RELEASE=unknown
ARG COMMIT=unknown
ARG BUILD_TIME=unknown

ENV RELEASE ${RELEASE}
ENV COMMIT ${COMMIT}
ENV BUILD_TIME ${BUILD_TIME}

RUN go build -v \
  		-ldflags "-s -w \
      -X ${PROJECT}/main.Release=${RELEASE} \
  		-X ${PROJECT}/main.Commit=${COMMIT} \
      -X ${PROJECT}/main.BuildTime=${BUILD_TIME}" \
  		-o batrium-udp2http-bridge ;

RUN find /go/src -type d

# final stage
FROM alpine:3

RUN apk update && apk upgrade
WORKDIR /app
COPY --from=build-env /go/src/github.com/liskl/batrium-udp2http-bridge/batrium-udp2http-bridge /app/
COPY --from=build-env /go/src/github.com/liskl/batrium-udp2http-bridge/static /app/static
COPY --from=build-env /go/src/github.com/liskl/batrium-udp2http-bridge/templates /app/templates

CMD ["/app/batrium-udp2http-bridge"]
