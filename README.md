[![Build Status](https://drone.julina.ch/api/badges/sybnex/drone-clair/status.svg)](https://drone.julina.ch/sybnex/drone-clair)

# drone-clair

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
