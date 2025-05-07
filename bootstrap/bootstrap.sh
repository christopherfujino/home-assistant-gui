#!/usr/bin/env bash

if [[ $(id -u) -ne 0 ]]; then
	echo "Must run as root"
	exit 1
fi

# D-Bus is required for bluetooth
docker run -d \
	--name homeassistant \
	--privileged \
	--restart=unless-stopped \
	-e TZ=America/Los_Angeles \
	-v "$PWD/config:/config" \
	-v /run/dbus:/run/dbus:ro \
	--network=host \
	ghcr.io/home-assistant/home-assistant:stable
