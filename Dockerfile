# Dockerfile
FROM alpine:3.14.2
COPY helloween \
	/usr/bin/helloween
ENTRYPOINT ["/usr/bin/helloween"]