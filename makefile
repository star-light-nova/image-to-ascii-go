install:
	# MacOS installation
	# HOMEBREW_NO_AUTO_UPDATE=1 - will skip auto-update
	# for every package that you have.
	HOMEBREW_NO_AUTO_UPDATE=1 brew install vips pkg-config
	pkg-config --cflags --libs vips
	go mod download && go mod verify

# Open a PR for other OS

# Docker
name=itaii
build:
	docker build -t $(name) .
dock-run:
	docker run --name $(name) --publish 8080:8080 $(name)
dock-stop:
	docker container stop $(name)
	docker container rm $(name)
dock-clean:
	docker container stop $(name) | \
	docker container rm $(name) | \
	docker image rm $(name)

# Default
run-dev:
	export CGO_CFLAGS_ALLOW=-Xpreprocessor && \
	go run main.go
run-prod:
	export CGO_CFLAGS_ALLOW=-Xpreprocessor && \
	export GIN_MODE=release && \
	go run main.go



