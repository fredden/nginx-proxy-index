# nginx proxy index

Intended to be used with [nginx-proxy](https://github.com/jwilder/nginx-proxy), this will produce a simple "landing page" listing connected websites.

For best results, we recommend setting `DEFAULT_HOST` environment variable to `localhost` in the nginx-proxy. For example:

```sh
docker run -d \
  -p 80:80 \
  -v /var/run/docker.sock:/tmp/docker.sock:ro \
  --net webproxy \
  -e DEFAULT_HOST=localhost \
  jwilder/nginx-proxy:alpine
```

## Installation

Clone this repository from GitHub:

```sh
git clone https://github.com/fredden/nginx-proxy-index.git
```

## Configuration

The following settings can be optionally added to a `.env` file in this directory.

* `DEFAULT_HOST` (default: `localhost`)
This is the hostname assigned to this website. We recommend setting this same variable on the nginx-proxy configuration too.

* `NETWORK` (default: `webproxy`)
The docker network where websites are discoverable.

## Start / Stop

```sh
# Start the service
docker-compose up -d

# Stop the service
docker-compose down -v
```
