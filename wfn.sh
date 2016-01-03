#!/bin/sh

echo "Starting wfn backend..."
wfn &

echo "Starting caddy frontend..."
caddy -conf /caddy/Caddyfile -root /caddy/public -quiet

echo "Caddy frontend exited..."
