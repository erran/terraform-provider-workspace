default: build

.PHONY: deps build test plan

deps:
	go get .

build:
	go install github.com/mitchellh/gox
	gox -ldflags "-X main.version=${VERSION}" \
			-osarch "darwin/amd64 darwin/arm64 linux/amd64 linux/arm64" \
			-output "build/{{.OS}}_{{.Arch}}/terraform-provider-workspace_$(VERSION)"

test:
	go test -v .
