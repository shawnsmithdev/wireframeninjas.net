#!/bin/sh
# TODO: Set it up so that this process dies if either child does.
ls
if [[ -x "/go/bin/wfn" ]]; then
    echo "Starting wfn backend..."
    /go/bin/wfn &
else
    echo "Could not find /go/bin/wfn"
fi

echo "In between"

if [[ -x "/go/bin/caddy" ]]; then
    echo "Starting caddy frontend..."
    /go/bin/caddy -root ./root
else
    echo "Could not find /go/bin/caddy"
fi
