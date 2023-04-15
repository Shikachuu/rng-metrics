all: container

container:
	buildah build -t ghcr.io/shikachuu/rng-metrics .