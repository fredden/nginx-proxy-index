`docker run -d -p 80:80 -v /var/run/docker.sock:/tmp/docker.sock:ro --net webproxy -e DEFAULT_HOST=localhost jwilder/nginx-proxy:alpine`
