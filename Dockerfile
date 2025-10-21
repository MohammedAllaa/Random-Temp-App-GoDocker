# Step 1: Use the latest Go base image
FROM golang:1.25-alpine

# Step 2: Set working directory
WORKDIR /app

# Step 3: Copy go.mod and go.sum (if exists)
COPY go.mod go.sum* ./

# Step 4: Download dependencies
RUN go mod download

# Step 5: Copy the entire project
COPY . .

# Step 6: Build the Go app
RUN go build -o main .

# Step 7: Run the app when container starts
CMD ["./main"]
