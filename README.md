# Webhook

a simple application that allows you to make a git pull after receiving a message from an external source, such as github, gitlab, bitbuket, etc.

# Docker-compose


```yaml
version: "3.3"

services:

  traefik:
    image: "traefik:v2.5"
    container_name: "traefik"
    command:
      - "--api.insecure=true"
      - "--providers.docker=true"
      - "--providers.docker.exposedbydefault=false"
      - "--entrypoints.web.address=:80"
    ports:
      - "80:80"
      - "8080:8080"
    volumes:
      - "/var/run/docker.sock:/var/run/docker.sock:ro"

  nginx:
    image: nginx:alpine
    restart: always
    volumes:
      - ./website:/usr/share/nginx/html
    labels:
      - "traefik.enable=true"
      - "traefik.http.routers.frontend.rule=Host(`nginx.localhost`)"
      - "traefik.http.routers.frontend.entrypoints=web"

  webhook:
    build:
      context: .
      dockerfile: Dockerfile
    volumes:
      # Mount app directory into /app
      - ./:/app
      # Mount .ssh into /root/.ssh
      - ~/.ssh:/root/.ssh
    environment:
      # Secret query string, like http://blabla.com/?secret=super-puper-secret
      - WEBHOOK_SECRET=super-puper-secret
      # Git repository directory /app/website
      - WEBHOOK_WORKDIR=./website
    expose:
      - 8000
    labels:
      - "traefik.enable=true"
      - "traefik.http.routers.webhook.rule=Host(`hook.localhost`)"
      - "traefik.http.routers.webhook.entrypoints=web"
      - "traefik.http.routers.webhook.service=webhook"
      - "traefik.http.services.webhook.loadbalancer.server.port=8000"
    ports:
      - 8000:8000
```