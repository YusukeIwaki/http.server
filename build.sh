#!/bin/sh
gox -osarch="darwin/amd64 linux/386 linux/amd64" -output="./out/bin/http.server_{{.OS}}_{{.Arch}}" -parallel=2
