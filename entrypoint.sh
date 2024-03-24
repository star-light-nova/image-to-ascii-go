#!/bin/bash

export CGO_CFLAGS_ALLOW="-Xpreprocessor"
echo "Flag for vips library set. $CGO_CFLAGS_ALLOW"

echo "Running Go server..."
./itaii
