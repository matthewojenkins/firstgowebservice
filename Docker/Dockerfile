FROM golang

ADD build/program /go/bin/programservice

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/programservice/program

# Document that the service listens on port 8080.
EXPOSE 8080
