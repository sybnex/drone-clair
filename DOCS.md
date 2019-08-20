---
date: 2016-01-01T00:00:00+00:00
title: Clair
author: jmccann
tags: [ docker, security ]
repo: jmccann/drone-clair
logo: clair.svg
image: jmccann/drone-clair
---

The Clair plugin submits your docker image to your [Clair](https://github.com/coreos/clair)
server to scan your docker image for security vulnerabilities.

The below pipeline configuration demonstrates simple usage:

```yaml
- name: scan-image
  image: sybex/drone-clair
  settings:
    url: https://clair.company.com
    username: <user>
    password: <pass>
    scan_image: debian:10
    security: Low
    threshold: 1
```

To verify https/ssl connections with a different CA certificate use `ca_cert`

```diff
- name: scan-image
  image: sybex/drone-clair
  settings:
    url: https://clair.company.com
    username: <user>
    password: <pass>
    scan_image: debian:10
    security: Low
    threshold: 1
+   ca_cert: |
+     -----BEGIN CERTIFICATE-----
+     MII...
+     -----END CERTIFICATE-----
```

# Secrets

The Clair plugin supports reading credentials from the Drone secret store. This is strongly recommended instead of storing credentials in the pipeline configuration in plain text.

```diff
- name: scan-image
  image: sybex/drone-clair
  settings:
    url: https://clair.company.com
-   username: <user>
-   password: <pass>
+   username:
+       from_secret: docker_username
+   password:
+       from_secret: docker_password
    scan_image: debian:10
    security: Low
    threshold: 1
```

Please see the Drone [documentation]({{< secret-link >}}) to learn more about secrets.

# Secret Reference

DOCKER_USERNAME
: paired with `username` - The username to authenticate to the docker registry with

DOCKER_PASSWORD
: paired with `password` - The password to authenticate to the docker registry with

CLAIR_URL
: paired with `url` - Clair server URL

CLAIR_OUTPUT
: paired with `security` - min. level for output (Low, Medium, High)

CLAIR_THRESHOLD
: paired with `threshold` - how many vuln. are acceptable

CLAIR_CA_CERT
: paired with `ca_cert` - The CA Cert to verify https with

# Parameter Reference

url
: Clair server URL

username
: Docker Registry username to download the `scan_image` from

password
: Docker Registry password to download the `scan_image` from

scan_image
: The docker image to scan.  Supports Docker Hub or private repos.

security
: The min. vuln. level to act on. 

threshold
: A value of hoa many vuln. are accepting before exit with >0

ca_cert
: The CA Cert to verify https with
