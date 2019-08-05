FROM golang:1.12.7-alpine3.10 AS build
RUN apk add --no-cache git make
COPY . /src/mattermost-housekeeper
WORKDIR /src/mattermost-housekeeper
RUN make linux

FROM alpine:3.10.1
COPY --from=build /src/mattermost-housekeeper/bin/mattermost-housekeeper-linux-amd64 /usr/bin/mattermost-housekeeper
ENTRYPOINT ["mattermost-housekeeper", "clean"]
