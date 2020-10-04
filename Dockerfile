# Get base Golang image to build the binary
FROM golang:1.15 as BuildImage

# Build binary
WORKDIR /build
COPY . .
RUN go build .

# Remove source and only keep binary build
FROM scratch

# Change dir & copy binary from build image
WORKDIR /go/bin
COPY --from=BuildImage /build/server .

# Expose port
EXPOSE 42069

# Launch server
ENTRYPOINT [ "./server" ]
