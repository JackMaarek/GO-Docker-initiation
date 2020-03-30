# Import go image
FROM golang:1.13-alpine
# Label for maintainer
LABEL maintainer="Jack Maarek, Alexandre Lellouche"
# Set the working directory inside the container
WORKDIR /go/src/app
# Copy the full project to currennt directory
COPY . .
# Run command to nstall the dependencies
RUN go mod download
# Build thego-app binary
RUN go build -o main .
# Port of for outside access
EXPOSE 8000
# Run the go-app executable
CMD ["./go/src/app/main"]