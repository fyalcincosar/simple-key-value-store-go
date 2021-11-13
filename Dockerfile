FROM golang:1.16-alpine
# Add a work directory
WORKDIR /in-memory-key-value-store
# Cache and install dependencies
COPY go.mod ./
RUN go mod download
# Copy app files
COPY . .
# Expose port
EXPOSE 8080
# Start app
RUN go build -o /main 
CMD [ "/main" ]


