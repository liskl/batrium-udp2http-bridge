module github.com/liskl/batrium-udp2http-bridge

go 1.16

require (
	github.com/go-kit/kit v0.12.0
	github.com/gorilla/mux v1.8.0
	github.com/prometheus/client_golang v1.11.0
	github.com/sirupsen/logrus v1.8.1
)

replace (
	github.com/miekg/dns => github.com/miekg/dns v1.1.43
	github.com/nats-io/jwt => github.com/nats-io/jwt/v2 v2.1.0
	github.com/nats-io/jwt/v2 => github.com/nats-io/jwt/v2 v2.1.0
)
