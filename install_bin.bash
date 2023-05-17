#!/bin/bash

if [[ "$OSTYPE" == "linux-gnu"* ]]; then
    os="linux"

    machine=$(uname -m)
    if [[ "$machine" == "aarch64" ]]; then
        arch="arm64"
    elif [[ "$machine" == "x86_64" ]]; then
        arch="x64"
    else
        echo "Unsupported architecture"
        exit 1
    fi

elif [[ "$OSTYPE" == "darwin"* ]]; then
    os="osx"

    arch=$(uname -m)
    if [[ "$arch" != "arm64" && "$arch" != "x86_64" ]]; then
        echo "Unsupported architecture"
        exit 1
    fi

else
    echo "Unsupported OS"
    exit 1
fi

# binディレクトリにlibvoicevox_core.dylibがあるか確認
if [  -e bin/libvoicevox_core.dylib ]; then
    echo "libvoicevox_core.dylib is already installed"
    exit 0
fi

binary="download-$os-$arch"

curl -sSfL https://github.com/VOICEVOX/voicevox_core/releases/latest/download/${binary} -o download
chmod +x download
./download

mv voicevox_core bin
rm download
