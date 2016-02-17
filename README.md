# Simple Reverse Proxy

An extremely simple reverse proxy using Go's standard library and configured by a single environment variable. Perfect for a minimal Docker container.

## Use

Provide `BACKEND_URL` environment variable, and simplerp will serve it on port 8080. Prebuilt minimal docker image available as `gwwfps/simplerp`.

## Build

Standard Go package building procedure applies, or run `docker run --rm -v "$(pwd):/src" -v /var/run/docker.sock:/var/run/docker.sock centurylink/golang-builder` to build the docker image.
