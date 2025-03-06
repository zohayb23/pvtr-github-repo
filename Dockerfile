FROM alpine:3.21 AS core
RUN apk add --no-cache wget tar unzip

WORKDIR /app
ARG VERSION=0.7.0
ARG PLATFORM=Linux_x86_64  # Change this based on your target system

RUN wget https://github.com/privateerproj/privateer/releases/download/v${VERSION}/privateer_${PLATFORM}.tar.gz
RUN tar -xzf privateer_${PLATFORM}.tar.gz

FROM golang:1.23.4-alpine3.21 AS plugin
RUN apk add --no-cache make git
WORKDIR /plugin
COPY . .
RUN make binary

FROM golang:1.23.4-alpine3.21
RUN apk add --no-cache make git && \
    mkdir -p /.privateer/bin
WORKDIR /.privateer/bin
COPY --from=core /app/privateer .
COPY --from=plugin /plugin/github-repo .
COPY --from=plugin /plugin/container-entrypoint.sh .

# The config file must be provided at run time.
# example: docker run -v /path/to/config.yml:/.privateer/bin/config.yml privateer-image
CMD ["./container-entrypoint.sh"]
