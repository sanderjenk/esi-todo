FROM golang:1.15.8 as builder
WORKDIR $GOPATH/src/github.com/codefresh-contrib/go-sample-app

COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

EXPOSE 8000
ENTRYPOINT ["go", "run", "main.go"]