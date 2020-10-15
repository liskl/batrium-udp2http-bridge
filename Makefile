APP?=batrium-udp2http-bridge
TAG?=latest
REGISTRY?=registry.infra.liskl.com

COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
RELEASE?=$(shell cat ./versionInfo)

clean:
	$(shell docker rm -f "${APP}" || true )
	$(shell docker rmi -f "${APP}:${TAG}" || true )

clean-ui:
	$(shell docker rm -f "${APP}-ui" || true )
	$(shell docker rmi -f "${APP}-ui:${TAG}" || true )

clean-tests:
	$(shell docker rm -f "${APP}-tests" || true )
	$(shell docker rm -f "${APP}-tests" || true )

format:
	go fmt ./;
	go fmt ./batrium;

build: clean format
	docker build \
	 --build-arg "BUILD_TIME=${BUILD_TIME}" \
	 --build-arg "COMMIT=${COMMIT}" \
	 --build-arg "RELEASE=${RELEASE}" \
	 -t "${REGISTRY}/${APP}:${TAG}" -f Dockerfile . ; \
	docker push "${REGISTRY}/${APP}:${TAG}";

build-ui: clean-ui
	docker build \
	 --build-arg "BUILD_TIME=${BUILD_TIME}" \
	 --build-arg "COMMIT=${COMMIT}" \
	 --build-arg "RELEASE=${RELEASE}" \
	 -t "${REGISTRY}/${APP}:${TAG}" -f Dockerfile.ui . ; \
	docker push "${REGISTRY}/${APP}-ui:${TAG}";


build-tests: clean-tests format
	docker build \
	 --build-arg "BUILD_TIME=${BUILD_TIME}" \
	 --build-arg "COMMIT=${COMMIT}" \
	 --build-arg "RELEASE=${RELEASE}" \
	 -t "${REGISTRY}/${APP}-client:${TAG}" -f Dockerfile.tests . ; \
	docker push "${REGISTRY}/${APP}-client:${TAG}";

build-all: build build-ui

run: build
	docker pull "${REGISTRY}/${APP}:${TAG}";
	docker run --rm --name "${APP}" -it \
		-p 18542:18542/udp \
		-p 8080:8080/tcp \
		"${REGISTRY}/${APP}:${TAG}" ;

run-ui: build-ui
	docker pull "${REGISTRY}/${APP}-ui:${TAG}";
	docker run --rm --name "${APP}-ui" -it \
		-p 8081:80/tcp \
		"${REGISTRY}/${APP}-ui:${TAG}" ;


test:
	 cd ./src/github.com/liskl/${APP} \
	 && clear; go test -v ./... \
	 && mkdir -p ./tests \
	 && go test -coverprofile tests/cp.out \
	 && go tool cover -html=tests/cp.out ;

release: test
	docker tag "${REGISTRY}/${APP}:${TAG}" "${REGISTRY}/${APP}:${RELEASE}";
	docker tag "${REGISTRY}/${APP}-ui:${TAG}" "${REGISTRY}/${APP}-ui:${RELEASE}";
	docker push "${REGISTRY}/${APP}:${RELEASE}";
	docker push "${REGISTRY}/${APP}-ui:${RELEASE}";
	bumpversion --current-version "${RELEASE}" --allow-dirty --commit patch versionInfo ;
