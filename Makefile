.DEFAULT_GOAL := build

BINARY = mattermost-housekeeper
REPOSITORY = github.com/anton-johansson/${BINARY}
IMAGE = antonjohansson/${BINARY}
VERSION = 0.0.0

GO_VERSION = $(shell go version | awk -F\go '{print $$3}' | awk '{print $$1}')
COMMIT = $(shell git rev-parse HEAD)
BUILD_DATE = $(shell date --utc +'%Y-%m-%dT%H:%M:%SZ')
PACKAGE_LIST = $$(go list ./...)
OUTPUT_DIRECTORY = ./bin
LDFLAGS = -ldflags " \
	-X ${REPOSITORY}/pkg/version.version=${VERSION} \
	-X ${REPOSITORY}/pkg/version.goVersion=${GO_VERSION} \
	-X ${REPOSITORY}/pkg/version.commit=${COMMIT} \
	-X ${REPOSITORY}/pkg/version.buildDate=${BUILD_DATE} \
	"

install:
	go get -v -d ./...

fmt:
	gofmt -s -d -e -w .

vet:
	go vet ${PACKAGE_LIST}

test: install
	go test ${PACKAGE_LIST}

linux: install
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${OUTPUT_DIRECTORY}/${BINARY}-linux-amd64 ./cmd/mattermost-housekeeper

darwin: install
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o ${OUTPUT_DIRECTORY}/${BINARY}-darwin-amd64 ./cmd/mattermost-housekeeper

windows: install
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o ${OUTPUT_DIRECTORY}/${BINARY}-windows-amd64.exe ./cmd/mattermost-housekeeper

build: linux darwin windows

docker:
	docker build -t ${IMAGE} .

docker-archive:
	docker save -o image.tar ${IMAGE}

docker-push:
	echo "${DOCKERHUB_PASSWORD}" | docker login -u "${DOCKERHUB_USERNAME}" --password-stdin
	docker tag ${IMAGE_NAME} ${IMAGE_NAME}:${VERSION}
	docker push ${IMAGE_NAME}:${VERSION}

clean:
	rm -rf ${OUTPUT_DIRECTORY}

