# Mattermost housekeeper

Mattermost Entpriese Edition provides a tool for retention policies that makes sure that older posts and attachments are automatically removed after a set period of time, to keep the disk usage low. Mattermost Team Edition does not have this utility and this tool is a workaround for that.


## Building

Binaries (requires `make` and `go`):

```shell
make linux
make darwin
make windows
make build # Builds all above
```

Docker (requires `make` and `docker`):

```shell
make docker
```


## Running

```shell
$ mattermost-housekeeper clean --data-dir /mattermost/data --database-host 192.168.123.123 --database-user mattermost --database-password s3cr3t --database-name mattermost
```
