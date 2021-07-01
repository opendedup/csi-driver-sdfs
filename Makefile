# Copyright 2017 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

CMDS=sdfsplugin
DEPLOY_FOLDER = ./deploy
CMDS=sdfsplugin
PKG = github.com/opendedup/csi-driver-sdfs
GINKGO_FLAGS = -ginkgo.v
GO111MODULE = on
GOPATH ?= $(shell go env GOPATH)
GOBIN ?= $(GOPATH)/bin
DOCKER_CLI_EXPERIMENTAL = enabled
export GOPATH GOBIN GO111MODULE DOCKER_CLI_EXPERIMENTAL

include release-tools/build.make

GIT_COMMIT = $(shell git rev-parse HEAD)
BUILD_DATE = $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
IMAGE_VERSION = $(shell git describe --tags)
LDFLAGS = -X ${PKG}/pkg/sdfs.driverVersion=${IMAGE_VERSION} -X ${PKG}/pkg/sdfs.gitCommit=${GIT_COMMIT} -X ${PKG}/pkg/sdfs.buildDate=${BUILD_DATE}
EXT_LDFLAGS = -s -w -extldflags "-static"
# Use a custom version for E2E tests if we are testing in CI
ifdef CI
ifndef PUBLISH
override IMAGE_VERSION := e2e-$(GIT_COMMIT)
endif
endif
IMAGENAME ?= sdfsplugin
REGISTRY ?= andyzhangx
REGISTRY_NAME ?= $(shell echo $(REGISTRY) | sed "s/.azurecr.io//g")
IMAGE_TAG = $(REGISTRY)/$(IMAGENAME):$(IMAGE_VERSION)
IMAGE_TAG_LATEST = $(REGISTRY)/$(IMAGENAME):latest

E2E_HELM_OPTIONS ?= --set image.sdfs.repository=$(REGISTRY)/$(IMAGENAME) --set image.sdfs.tag=$(IMAGE_VERSION) --set image.sdfs.pullPolicy=Always
E2E_HELM_OPTIONS += ${EXTRA_HELM_OPTIONS}

all: sdfs

.PHONY: verify
verify: unit-test
	hack/verify-all.sh

.PHONY: unit-test
unit-test:
	go test -covermode=count -coverprofile=profile.cov ./pkg/... -v

.PHONY: sanity-test
sanity-test: sdfs
	./test/sanity/run-test.sh

.PHONY: integration-test
integration-test: sdfs
	./test/integration/run-test.sh

.PHONY: local-build-push
local-build-push: sdfs
	docker build -t $(LOCAL_USER)/sdfsplugin:latest .
	docker push $(LOCAL_USER)/sdfsplugin

.PHONY: local-k8s-install
local-k8s-install:
	echo "Instlling locally"
	kubectl apply -f $(DEPLOY_FOLDER)/rbac-csi-sdfs-controller.yaml
	kubectl apply -f $(DEPLOY_FOLDER)/csi-sdfs-driverinfo.yaml
	kubectl apply -f $(DEPLOY_FOLDER)/csi-sdfs-controller.yaml
	kubectl apply -f $(DEPLOY_FOLDER)/csi-sdfs-node.yaml
	echo "Successfully installed"

.PHONY: local-k8s-uninstall
local-k8s-uninstall:
	echo "Uninstalling driver"
	kubectl delete -f $(DEPLOY_FOLDER)/csi-sdfs-controller.yaml --ignore-not-found
	kubectl delete -f $(DEPLOY_FOLDER)/csi-sdfs-node.yaml --ignore-not-found
	kubectl delete -f $(DEPLOY_FOLDER)/csi-sdfs-driverinfo.yaml --ignore-not-found
	kubectl delete -f $(DEPLOY_FOLDER)/rbac-csi-sdfs-controller.yaml --ignore-not-found
	echo "Uninstalled SDFS driver"

.PHONY: sdfs
sdfs:
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags "${LDFLAGS} ${EXT_LDFLAGS}" -mod vendor -o bin/sdfsplugin ./cmd/sdfsplugin

.PHONY: container
container: sdfs
	docker build --no-cache -t $(IMAGE_TAG) .

.PHONY: push
push:
	docker push $(IMAGE_TAG)

.PHONY: push-latest
push-latest:
	docker tag $(IMAGE_TAG) $(IMAGE_TAG_LATEST)
	docker push $(IMAGE_TAG_LATEST)

.PHONY: install-sdfs-server
install-sdfs-server:
	kubectl apply -f ./deploy/example/sdfs-provisioner/sdfs-server.yaml

.PHONY: install-helm
install-helm:
	curl https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3 | bash

.PHONY: e2e-bootstrap
e2e-bootstrap: install-helm
	docker pull $(IMAGE_TAG) || make container push
	helm install csi-driver-sdfs ./charts/latest/csi-driver-sdfs --namespace kube-system --wait --timeout=15m -v=5 --debug \
		${E2E_HELM_OPTIONS} \
		--set controller.logLevel=8 \
		--set node.logLevel=8

.PHONY: e2e-teardown
e2e-teardown:
	helm delete csi-driver-sdfs --namespace kube-system

.PHONY: e2e-test
e2e-test:
	if [ ! -z "$(EXTERNAL_E2E_TEST)" ]; then \
		bash ./test/external-e2e/run.sh;\
	else \
		go test -v -timeout=0 ./test/e2e ${GINKGO_FLAGS};\
	fi
