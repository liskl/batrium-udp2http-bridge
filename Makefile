APP?=batrium-udp2http-bridge
TAG?=latest
REGISTRY?=registry.infra.liskl.com

COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
RELEASE?=$(shell cat ./versionInfo)

clean:
	$(shell docker rm -f "${APP}" "${APP}-tests" || true )
	$(shell docker rmi -f "${APP}:${TAG}" || true )

format:
	go fmt ;

build: clean format
	docker build \
	 --build-arg "BUILD_TIME=${BUILD_TIME}" \
	 --build-arg "COMMIT=${COMMIT}" \
	 --build-arg "RELEASE=${RELEASE}" \
	 -t "${REGISTRY}/${APP}:${TAG}" . ; \
	docker push "${REGISTRY}/${APP}:${TAG}";

build-client: clean format
	docker build \
	 --build-arg "BUILD_TIME=${BUILD_TIME}" \
	 --build-arg "COMMIT=${COMMIT}" \
	 --build-arg "RELEASE=${RELEASE}" \
	 -t "${REGISTRY}/${APP}-client:${TAG}" -f Dockerfile.client . ; \
	docker push "${REGISTRY}/${APP}-client:${TAG}";

run: build
	docker pull "${REGISTRY}/${APP}:${TAG}";
	docker run --rm --name "${APP}" -it \
		-p 18542:18542/udp \
		-p 8080:8080/tcp \
		"${REGISTRY}/${APP}:${TAG}" ;

test:
	 cd ./src/github.com/liskl/${APP} \
	 && clear; go test -v ./... \
	 && mkdir -p ./tests \
	 && go test -coverprofile tests/cp.out \
	 && go tool cover -html=tests/cp.out ;

release: test
	docker tag "${REGISTRY}/${APP}:${TAG}" "${REGISTRY}/${APP}:${RELEASE}";
	docker push "${REGISTRY}/${APP}:${RELEASE}";
	bumpversion --current-version "${RELEASE}" --allow-dirty --commit patch versionInfo ;
