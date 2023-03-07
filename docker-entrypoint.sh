#!/bin/sh
set -e

if [ $(echo "$1" | cut -c1) = "-" ]; then
    echo "$0: assuming arguments for mdnscli service"
    set -- mdnscli service "$@"
fi

if [ "$1" = "mdnscli"  ]; then
    echo "$0: assuming arguments for mdnscli service"
    set -- "$@"
fi

if [ "$1" = "service" ]; then
    echo "$0: assuming arguments for mdnscli service"
    set -- mdnscli "$@"
fi

if [ "$1" = "client" ]; then
    echo "$0: assuming arguments for mdnscli client"
    set -- mdnscli "$@"
fi

echo "run some: $@"
exec "$@"
