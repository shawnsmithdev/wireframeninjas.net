#!/bin/sh
# TODO: Set it up so that this process dies if either child does.
echo "Starting wfn backend..."
/go/bin/wfn &
echo "Starting caddy frontend..."
/go/bin/caddy -root ./root
