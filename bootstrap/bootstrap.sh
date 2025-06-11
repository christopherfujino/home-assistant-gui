#!/usr/bin/env bash

set -euo pipefail

ROOT="$(dirname "$(dirname "$(realpath "$BASH_SOURCE")")")"

if [[ $(id -u) -eq 0 ]]; then
  echo "Must not run as root"
  exit 1
fi

if [[ ! -d "$ROOT/ignore" ]]; then
  mkdir "$ROOT/ignore"
fi

# Ensure fonts are cached
LOCAL_FONT_PATH="$ROOT/ignore/font.zip"
if [[ ! -f "$LOCAL_FONT_PATH" ]]; then
  curl -L 'https://style64.org/file/C64_TrueType_v1.2.1-STYLE.zip' -o "$LOCAL_FONT_PATH"
  unzip "$LOCAL_FONT_PATH" -d "$ROOT/ignore"
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
