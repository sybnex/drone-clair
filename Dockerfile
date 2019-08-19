# Docker stage build if needed or wanted
#FROM golang
#WORKDIR /go/src/github.com/sybex/drone-clair
#RUN go get github.com/urfave/cli && \
#    go get github.com/Sirupsen/logrus
#COPY main.go .
#COPY plugin.go .
#RUN CGO_ENABLED=0 go build -ldflags "-s -w -X main.revision=$(git rev-parse HEAD)" -a -o drone-clair

FROM alpine
ADD https://github.com/optiopay/klar/releases/download/v2.4.0/klar-2.4.0-linux-amd64 /usr/local/bin/klar
RUN apk --no-cache add curl ca-certificates && \
    chmod 0755 /usr/local/bin/klar

ADD drone-clair /bin/
#COPY --from=0 /go/src/github.com/sybex/drone-clair /bin/

ENTRYPOINT ["/bin/drone-clair"]
