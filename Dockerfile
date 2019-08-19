# Docker image for the Drone clair plugin
#
#     cd $GOPATH/src/github.com/jmccann/drone-clair
#     go build
#     docker build --rm=true -t jmccann/drone-clair .

FROM alpine

ADD https://github.com/optiopay/klar/releases/download/v2.4.0/klar-2.4.0-linux-amd64 /usr/local/bin/klar

RUN apk --no-cache add curl ca-certificates && \
    chmod 0755 /usr/local/bin/klar

ADD drone-clair /bin/
ENTRYPOINT ["/bin/drone-clair"]
