FROM alpine:3.2

RUN apk add --update tar curl

RUN curl "https://caddyserver.com/download/build?os=linux&arch=amd64&features=prometheus" \
    | tar --no-same-owner -C /usr/bin/ -xz caddy

COPY Caddyfile /etc/Caddyfile
COPY public /www

EXPOSE 80

ENTRYPOINT ["/usr/bin/caddy"]
CMD ["--conf", "/etc/Caddyfile"]