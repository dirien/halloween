# Dockerfile
FROM alpine:3.14.2
COPY halloween \
	/usr/bin/halloween
ENTRYPOINT ["/usr/bin/halloween"]