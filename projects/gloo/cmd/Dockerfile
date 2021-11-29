# ARG ENVOY_IMAGE

# FROM $ENVOY_IMAGE
FROM quay.io/solo-io/envoy-gloo:1.19.0-patch3
ARG GOARCH=amd64

RUN apk upgrade --update-cache \
    && apk add ca-certificates \
    && rm -rf /var/cache/apk/*

COPY gloo-linux-$GOARCH /usr/local/bin/gloo

USER 10101

ENTRYPOINT ["/usr/local/bin/gloo"]