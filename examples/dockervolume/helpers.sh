#!/usr/bin/env bash
set -euo pipefail

create_docker_volume() {
  # Create Docker volume if it doesn't already exist
  docker volume ls | grep myvol || \
    docker volume create myvol

  # Add files to volume via a helper container
  docker create -v myvol:/data --name helper busybox true
  # Could copy the entire app directory in
  # docker cp /cnab/app helper:/data
  docker cp /cnab/app/echo.txt helper:/data
  docker cp /cnab/app/config.toml helper:/data
  docker rm helper > /dev/null 2>&1
}

update_docker_volume() {
  create_docker_volume
}

remove_docker_volume() {
  docker volume rm myvol
}

print_app_logs() {
  docker run -v myvol:/data --name logreader busybox cat /data/log.txt
  docker rm logreader > /dev/null 2>&1
}

# Call the requested function and pass the arguments as-is
"$@"