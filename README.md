# Bodo

![bodo](https://github.com/user-attachments/assets/57fdbfb4-903d-44c3-b735-26af5866c9bd)

## About

Bodo is a simple dashboard that can be used as dashboard for self hosted
services or as an browser homepage with links to your favorite services.

### Features

* simple & fast
* multiple themes
* easily configurable and customizable

## Instalation/deployment

### Manual

1. Clone git repository `git clone https://github.com/piotrkira/bodo.git`
2. Install bodo with `cd bodo && make install`
3. Run `bodo`

I recommended you to run Bodo in background as systemd service:

4. Create Systemd user service: `$HOME/.config/systemd/user/bodo.service`

```
[Unit]
Description=Bodo Dashboard

[Service]
Type=simple
StandardOutput=journal
ExecStart=/usr/local/bin/bodo

[Install]
WantedBy=default.target
```

5. Enable and start service `systemctl --user --now enable bodo.service`
6. Check if everything is okay `systemctl --user status bodo.service`

### Deployment with docker

```sh
docker run --name bodo \
    -v ./data:/data:ro \
    -p 8080:8080 \
    -d \
    ghcr.io/piotrkira/bodo:latest \
    --config /data/config.yaml \
    --themes /data/themes.yaml \  # optional if you want to include custom thems.yaml
```

### Deployment with docker compose

```yaml
services:
  bodo:
    image: ghcr.io/piotrkira/bodo:latest
    ports:
      - "8080:8080"
    command: --config /data/config.yaml
    volumes:
     - ./data:/data:ro
```

## Configuration

To configure bodo edit config.yaml file, by default located in `/etc/bodo/config.yaml`

Example configuration:

```yaml
version: 1
title: "Bodo Dashboard"                 # Dashboard title
columns: 3                              # Number of columns
theme: cattpucin_mocha                  # Theme name, choose from themes available in themes.yaml
font: Arial                             # Font name to use
sections:                               # List of dashboard sections
  - name: Media                         # Section name
    entries:                            # List of services
      - name: Jellyfin                  # Name of the serevice
        url: http://192.168.0.83:8096   # URL to service
      - name: Navidrone
        url: https://navidrome.home.local
  - name: Smart Home
    entries:
      - name: Homeassistant
        url: http://192.168.0.10
  - name: Favorite blogs
    entries:
      - name: Ratfactor
        url: https://ratfactor.com/
```

## Theming

Creating custom themes is quite easy, you just have to edit themes.yaml file,
add new entry, configure color and change theme name in config.yaml.

Example theme:

```yaml
retro:
  text_color: "#cd00cd"
  background_color: "#000000"
  primary_color: "#cdcd00"
```

## Contributing

For new features create PR or open new Issue.

## License

[MIT](https://github.com/piotrkira/bodo/blob/main/LICENSE)
