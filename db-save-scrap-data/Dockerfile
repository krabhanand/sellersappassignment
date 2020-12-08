FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /home/kumar/go/sellersapp/sellersappassignment/db-save-scrap-data

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o main .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist/b

# Copy binary from build to main folder
RUN cp /home/kumar/go/sellersapp/sellersappassignment/db-save-scrap-data/main .

# Export necessary port
EXPOSE 8888

# Command to run when starting the container
CMD ["/dist/b/main"]