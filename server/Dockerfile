# Get base Golang image to build the binary
FROM golang:1.15 as BuildImage

WORKDIR /build
COPY . .
RUN go build .

# Remove source and only keep binary build
FROM scratch

WORKDIR /go/bin
COPY --from=BuildImage /build/server .

EXPOSE 42069

ENTRYPOINT [ "./server" ]
