FROM golang:1.12.7-alpine3.10 AS build
RUN apk add --no-cache git make
COPY . /src/mattermost-maid
WORKDIR /src/mattermost-maid
RUN make linux

FROM alpine:3.10.1
COPY --from=build /src/mattermost-maid/bin/mattermost-maid-linux-amd64 /usr/bin/mattermost-maid
ENTRYPOINT ["mattermost-maid", "clean"]
