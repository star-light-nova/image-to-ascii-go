#!/bin/bash

if [[ "$1" -eq 1 && $1 == "install" ]]; then
    # MacOS installation
    # HOMEBREW_NO_AUTO_UPDATE=1 - will skip auto-update
    # for every package that you have.
    HOMEBREW_NO_AUTO_UPDATE=1 brew install vips pkg-config
    pkg-config --cflags --libs vips
fi
# Feel free to add installation for other OS'
export CGO_CFLAGS_ALLOW="-Xpreprocessor"
echo "Flag for vips library set. $CGO_CFLAGS_ALLOW"

echo "Running Go server..."
go run main.go

