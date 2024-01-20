# Use a minimal base image.
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the previous stage (build stage)
COPY myapp /app/myapp

# Expose the port on which your application listens
EXPOSE 3030

# Define the command to run your application
CMD ["./myapp"]

