FROM golang:latest AS build
WORKDIR /go/src/github.com/loqutus/O-1
COPY . ./
RUN make

FROM alpine:latest
WORKDIR /
COPY --from=build /go/src/github.com/loqutus/O-1/bin/o1-linux /usr/bin/o1
CMD ["/usr/bin/o1"]
EXPOSE 6969