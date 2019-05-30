FROM golang:alpine AS builder

MAINTAINER Qasim "qasim@zaidi.me"

LABEL version="1.0"

RUN apk update; apk add git
RUN go get -u github.com/golang/dep/cmd/dep

ARG repo
ENV PACKAGE_PATH "$GOPATH/src/${repo}"
WORKDIR "$PACKAGE_PATH"
RUN mkdir -p $PACKAGE_PATH

COPY Gopkg.toml Gopkg.lock $PACKAGE_PATH/
RUN dep ensure --vendor-only -v
COPY . ./

ARG binary_name
WORKDIR "$PACKAGE_PATH/cmd/${binary_name}"

RUN go build -o $GOPATH/bin/${binary_name} -ldflags "-X main.Version=0.0.1" && cp $GOPATH/bin/${binary_name} /${binary_name}

FROM alpine:latest
WORKDIR /root/

#copy the executable
COPY --from=builder /gocover .

# Expose emitter ports
EXPOSE 9000

# Start the broker
ENTRYPOINT ["./gocover" ]
