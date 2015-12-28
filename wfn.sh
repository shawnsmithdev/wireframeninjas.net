#!/bin/sh
echo "Starting wfn backend..."
/go/bin/wfn &
echo "Starting caddy frontend..."
/go/bin/caddy -root ./root
