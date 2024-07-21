FROM golang:1.22

WORKDIR /go/src

ENV PATH="/go/bin:${PATH}"
ENV GO111MODULE=on

COPY . .

# Install/update necessary libraries
RUN go get github.com/spf13/cobra@latest && \
    go install go.uber.org/mock/mockgen@latest && \
    go install github.com/spf13/cobra-cli@latest

# Install sqlite3
RUN apt-get update && apt-get install -y sqlite3

# Change UID of www-data user to 1000
RUN usermod -u 1000 www-data

# Create cache directory and set permissions
RUN mkdir -p /var/www/.cache
RUN chown -R www-data:www-data /go
RUN chown -R www-data:www-data /var/www/.cache

# Switch to www-data user
USER www-data

# Command to keep container running (useful for development)
CMD ["tail", "-f", "/dev/null"]
