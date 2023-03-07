#!/bin/sh
set -e

if [ $(echo "$1" | cut -c1) = "-" ]; then
    echo "$0: assuming arguments for mdnscli"
    set -- mdnscli "$@"
fi

echo "run some: $@"
exec "$@"
