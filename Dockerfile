FROM golang:1.19


COPY --from=rclone/rclone:1.50.2 /usr/local/bin/rclone /usr/local/bin/
