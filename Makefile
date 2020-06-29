all: build-image e2e

build-image: gofmt clean cross
	TAG=$(USER)-dev ./hack/build-image.sh

e2e:
	TAG=$(USER)-dev ./hack/e2e.sh

clean:
	rm -f $(name)

# Checks if Google criteria for releasing code as OSS are met
cross:
	if [ -e /google/data/ro/teams/opensource/cross ]; then /google/data/ro/teams/opensource/cross .; fi

# Formats go source code with gofmt
gofmt:
	find . -type f -name '*.go' | grep -v '/vendor/' | xargs gofmt -w

.PHONY: all build-image clean cross e2e gofmt
