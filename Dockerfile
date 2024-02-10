FROM golang:1.23 AS build

# Move current project to a valid go path
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# https://docs.docker.com/engine/reference/builder/#copy
# Ensure you have a .dockerignore file to exclude sensitive files/directories
# Copy the source code (be specific about what to include)
COPY main.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

# Final stage: Use a distroless image for running the app
FROM gcr.io/distroless/base-debian10:nonroot

# Copy the binary from the build stage
COPY --from=build /app/myapp /app/myapp

# Non-root user setup
USER nonroot

# Run app
EXPOSE 8080
ENTRYPOINT ["/app/myapp"]
