#!/bin/sh

VERSION=0.1.0

mkdir -p dist

TARGETS="
	darwin-amd64
	freebsd-386
	freebsd-amd64
	freebsd-arm
	linux-386
	linux-amd64
	linux-arm
	linux-arm64
	linux-ppc64
	openbsd-386
	openbsd-amd64
	openbsd-arm
"

# Remove empty lines and indentation
TARGETS=$(echo "$TARGETS" | sed -e '/^$/d' -e 's/	//')

for target in $TARGETS; do
	os=$(echo "$target" | cut -d'-' -f 1 -)
	arch=$(echo "$target" | cut -d'-' -f 2 -)

	GOOS="$os" \
	GOARCH="$arch" \
		go build -v

	tar cjvf "timetasker_${VERSION}_${os}_${arch}.tar.bz2" timetasker
	mv "timetasker_${VERSION}_${os}_${arch}.tar.bz2" dist
done
