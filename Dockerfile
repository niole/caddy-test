FROM caddy:2.9.1-builder AS builder
COPY . .
COPY site /srv
RUN xcaddy build --with github.com/niole/test-caddy-module=.

FROM caddy:2.9.1

COPY --from=builder /usr/bin/caddy /usr/bin/caddy

COPY Caddyfile /etc/caddy/Caddyfile
COPY site /srv
