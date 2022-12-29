FROM golang:1.19


COPY --from=rclone/rclone:1.61 /usr/local/bin/rclone /usr/local/bin/
COPY . /app

WORKDIR /app
RUN go mod tidy
