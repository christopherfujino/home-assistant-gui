#!/usr/bin/env bash

set -euo pipefail

if [[ $(id -u) -eq 0 ]]; then
  echo "Must not run as root"
  exit 1
fi

if [[ ! -d '../ignore' ]]; then
  mkdir ../ignore
fi

LOCAL_FONT_PATH='../ignore/font.zip'
if [[ ! -f "$LOCAL_FONT_PATH" ]]; then
  curl -L 'https://style64.org/file/C64_TrueType_v1.2.1-STYLE.zip' -o "$LOCAL_FONT_PATH"
  unzip "$LOCAL_FONT_PATH" -d '../ignore'
fi

# D-Bus is required for bluetooth
sudo docker run -d \
  --name homeassistant \
  --privileged \
  --restart=unless-stopped \
  -e TZ=America/Los_Angeles \
  -v "$PWD/config:/config" \
  -v /run/dbus:/run/dbus:ro \
  --network=host \
  ghcr.io/home-assistant/home-assistant:stable
