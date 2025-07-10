# Planttadore UI

A minimal GUI frontend for a portable
[Home Assistant](https://www.home-assistant.io/) install.

Designed to be run on the same device that the Home Assistant is running on.

## Dependencies

Install the base dependencies required to run the application:

```bash
sudo apt install \
     libgl1-mesa-dev \
     libxi-dev \
     libxcursor-dev \
     libxrandr-dev \
     libxinerama-dev \
     libwayland-dev \
     libxkbcommon-dev
```

Install the go compiler toolchain to be able to compile the app:

```bash
sudo apt install golang
```

And Docker to be able to run the Home Assistant within a container:

```bash
sudo apt install docker.io
```

## Setup

Run the `bootstrap.sh` to download the third-party font and the home assistant
Docker image:

```bash
./bootstrap/bootstrap.sh
```

You will need to populate your own configuration file, named `config.json`.
You can start from the example file in this repo:

```bash
cp ./config.json.example config.json
# Fill in your home-assistant.io token and path to server
nvim ./config.json
```

You will need to fill in your own `TOKEN` and `SENSOR_NAMES` values.
