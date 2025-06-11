A minimal GUI frontend for a portable
[Home Assistant](https://www.home-assistant.io/) install.

Designed to be run on the same device that the Home Assistant is running on.

## Dependencies

```bash
sudo apt install libgl1-mesa-dev libxi-dev libxcursor-dev libxrandr-dev libxinerama-dev libwayland-dev libxkbcommon-dev
```

## Setup

```bash
./bootstrap/bootstrap.sh
```

```bash
cp ./config.json.example config.json
# Fill in your home-assistant.io token and path to server
nvim ./config.json
```
