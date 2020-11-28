FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/matthewojenkins/firstgowebservice

RUN go build github.com/matthewojenkins/firstgowebservice

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/firstgowebservice

# Document that the service listens on port 8080.
EXPOSE 8080