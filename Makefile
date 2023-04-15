all: container

container:
	buildah build -t ghcr.io/shikachuu/random-metrics-generator .