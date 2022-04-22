# Copyright 2021 The batrium-udp2http-bridge Authors.
# SPDX-License-Identifier: Apache-2.0
#
# Makefile for batrium-udp2http-bridge service.

SHELL := /usr/bin/env bash
GOOS = $(shell go env GOOS)
GOARCH = $(shell go env GOARCH)
MYGOBIN = $(shell go env GOBIN)
ifeq ($(MYGOBIN),)
MYGOBIN = $(shell go env GOPATH)/bin
endif
export PATH := $(MYGOBIN):$(PATH)

APP?=batrium-udp2http-bridge
TAG?=latest
REGISTRY?=liskl

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
	go fmt ./client;

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
	 -t "${REGISTRY}/${APP}-ui:${TAG}" -f Dockerfile.ui . ; \
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
		-p 9001:9001/tcp \
		"${REGISTRY}/${APP}:${TAG}" ;

run-ui: build-ui
	docker pull "${REGISTRY}/${APP}-ui:${TAG}";
	docker run --rm --name "${APP}-ui" -it \
		-p 8081:80/tcp \
		"${REGISTRY}/${APP}-ui:${TAG}" ;


test:
	go test -v ./... \
	&& mkdir -p ./tests \
	&& go test -coverprofile tests/cp.out \
	&& go tool cover -html=tests/cp.out ;

snyk: build-ui
	snyk test \
	&& snyk test --docker "${REGISTRY}/${APP}-ui:${TAG}" --file="Dockerfile" \
	&& snyk test --docker "${REGISTRY}/${APP}:${TAG}" --file="Dockerfile" ;

release: test
	docker tag "${REGISTRY}/${APP}:${TAG}" "${REGISTRY}/${APP}:${RELEASE}";
	docker tag "${REGISTRY}/${APP}-ui:${TAG}" "${REGISTRY}/${APP}-ui:${RELEASE}";
	docker push "${REGISTRY}/${APP}:${RELEASE}";
	docker push "${REGISTRY}/${APP}-ui:${RELEASE}";
	bumpversion --current-version "${RELEASE}" --allow-dirty --commit patch versionInfo ;
