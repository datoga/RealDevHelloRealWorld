# Use your favorite image to create your HTTP server,
# like django, flask, express, or even an nginx!

# Docker stack is for extreme flexibility and control,
# hope you enjoy it!

# README.md provides a solution using nginx, if you wanna check it out.

FROM golang AS builder

ENV GO111MODULE=on

RUN mkdir -p /go/src/HelloRealWorld
WORKDIR /go/src/HelloRealWorld

ADD . /go/src/HelloRealWorld

WORKDIR /go/src/HelloRealWorld

RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -ldflags '-w -s -extldflags "-static"' .

FROM alpine

RUN mkdir /app 
WORKDIR /app
COPY --from=builder /go/src/HelloRealWorld/RealDevHelloRealWorld .

EXPOSE 8080

CMD ["/app/RealDevHelloRealWorld"]
