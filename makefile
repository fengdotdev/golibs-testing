# for updating the version of the project and pushing the tag to the repository
TESTING_VERSION = 0.0.5

updatev:
		git tag v${TESTING_VERSION} && git push origin v${TESTING_VERSION}

test:
	go test -v ./...
