FROM jwilder/docker-gen

RUN mkdir -p /app/pub
COPY index.html /app/pub/
COPY template.go /app/

CMD ["-watch", "/app/template.go", "/app/pub/index.html"]
