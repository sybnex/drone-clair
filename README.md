# drone-clair [![Build Status](https://drone.julina.ch/api/badges/sybnex/drone-clair/status.svg)](https://drone.julina.ch/sybnex/drone-clair) [![Pull Count](https://badgen.net/docker/pulls/sybex/drone-clair)](https://hub.docker.com/r/sybex/drone-clair)

Drone plugin to scan docker images with [Clair](https://github.com/coreos/clair).

For the usage information and a listing of the available options please take a look at [the docs](DOCS.md).

## Build

Build the binary with the following commands:

```
go build
go test
```

## Usage

Execute from the working directory:

```
docker run --rm \
  sybex/drone-clair --url http://clair.company.com --username johndoe --password mysecret \
                    --security Low --threshold 1 --scan_image python:2.7
```

Using from drone v1:

```
- name: scan-image
  image: sybex/drone-clair
  settings:
    url:
      from_secret: clair_url
    username:
      from_secret: docker_username
    password:
      from_secret: docker_password
    scan_image: sybex/drone-clair
    security: Low
    threshold: 1
```
