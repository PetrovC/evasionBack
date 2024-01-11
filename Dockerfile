FROM golang:latest AS build

# Move current project to a valid go path
COPY . /evasion
WORKDIR /evasion

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files have not changed
RUN go mod download

# Install Revel CLI
RUN go get github.com/revel/cmd/revel

# Install Revel CLI using go install (replace the version with the one you need)
RUN go install github.com/revel/cmd/revel@latest

# Ensure $GOPATH/bin is in your path for the Revel CLI to be found
ENV PATH="${PATH}:/root/go/bin"

# Run app in production mode
EXPOSE 9000
ENTRYPOINT revel run -a . /evasion

