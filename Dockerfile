# Golang base image for builder
FROM golang:1.13 AS builder

# Set builder working directory
WORKDIR $GOPATH/src/eDOT-test

# Copy repository to builder
COPY . .

# Build the Golang app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o /eDOT-test

# Alpine Linux base image for Golang app
FROM alpine:3.13

# Labels
LABEL Name="eDOT-test"
LABEL Version="1.0"

# Set working directory
WORKDIR /root/

# Install packages
RUN apk add --update-cache --no-cache ca-certificates tzdata util-linux

# Create directories
RUN mkdir conf && mkdir logs
# Copy the app.conf file from local
COPY conf/app.conf /root/conf/app.conf


# Copy only the binary file from builder
COPY --from=builder /eDOT-test /root/

# Listen to service port and admin/management port
EXPOSE 9900/tcp 9901/tcp

# Run the Golang app
CMD ["/root/eDOT-test"]
