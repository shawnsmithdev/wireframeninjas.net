#!/bin/sh

echo "Starting wfn backend..."
wfn &

echo "Starting caddy frontend..."
caddy -agree -email $CADDY_EMAIL -conf /caddy/Caddyfile -root /caddy/public

echo "Caddy frontend exited..."
