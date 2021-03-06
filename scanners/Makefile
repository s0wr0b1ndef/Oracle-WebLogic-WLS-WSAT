BINARY=CVE-2017-10271
GOARCH=amd64
HASH=$(shell git rev-parse HEAD)
BUILDDATE=$(shell date -u '+%Y-%m-%dT%k:%M:%SZ')
GOVERSION=$(shell go version | sed -e 's/ /|^|/g')
VERSION=1.5.1
LDFLAGS=-ldflags "-s -X github.com/kkirsche/$(BINARY)/scanners/cmd.BuildHash=$(HASH) -X github.com/kkirsche/$(BINARY)/scanners/cmd.BuildTime=$(BUILDDATE) -X github.com/kkirsche/$(BINARY)/scanners/cmd.BuildVersion=$(VERSION) -X github.com/kkirsche/$(BINARY)/scanners/cmd.BuildGoVersion=$(GOVERSION)"

lint:
	golint ./...

vet:
	go vet ./...

clean:
	rm -rf bin

install:
	go install -race -v

binary-depends:
	mkdir -p bin

# Builds
darwin-build: vet lint binary-depends
	env GOOS=darwin GOARCH=$(GOARCH) go build $(LDFLAGS) -race -v -o bin/$(BINARY).release.$(VERSION).$(GOARCH).darwin

dragonfly-build: vet lint binary-depends
	env GOOS=dragonfly GOARCH=$(GOARCH) go build $(LDFLAGS) -v -o bin/$(BINARY).release.$(VERSION).$(GOARCH).dragonfly

freebsd-build: vet lint binary-depends
	env GOOS=freebsd GOARCH=$(GOARCH) go build $(LDFLAGS) -v -o bin/$(BINARY).release.$(VERSION).$(GOARCH).freebsd

linux-build: vet lint binary-depends
	env GOOS=linux GOARCH=$(GOARCH) go build $(LDFLAGS) -v -o bin/$(BINARY).release.$(VERSION).$(GOARCH).linux

netbsd-build: vet lint binary-depends
	env GOOS=netbsd GOARCH=$(GOARCH) go build $(LDFLAGS) -v -o bin/$(BINARY).release.$(VERSION).$(GOARCH).netbsd

openbsd-build: vet lint binary-depends
	env GOOS=openbsd GOARCH=$(GOARCH) go build $(LDFLAGS) -v -o bin/$(BINARY).release.$(VERSION).$(GOARCH).openbsd

solaris-build: vet lint binary-depends
	env GOOS=solaris GOARCH=$(GOARCH) go build $(LDFLAGS) -v -o bin/$(BINARY).release.$(VERSION).$(GOARCH).solaris

windows-build: vet lint binary-depends
	env GOOS=windows GOARCH=$(GOARCH) go build $(LDFLAGS) -v -o bin/$(BINARY).release.$(VERSION).$(GOARCH).windows.exe

build: darwin-build dragonfly-build freebsd-build linux-build netbsd-build openbsd-build solaris-build windows-build

.PHONY: vet install binary-depends lint
