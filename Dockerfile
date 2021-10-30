# Dockerfile
FROM alpine:3.14.2
COPY minectl \
	/usr/bin/helloween
ENTRYPOINT ["/usr/bin/helloween"]