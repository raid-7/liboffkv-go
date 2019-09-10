#!/bin/bash

if [[ -z "$BOOT_DESTINATION" ]]; then
    BOOT_DESTINATION="$HOME"
fi

TRIPLET="$1"_amd64
EXT=zip
curl -L https://releases.hashicorp.com/consul/1.5.1/consul_1.5.1_${TRIPLET}.${EXT} -o "$BOOT_DESTINATION"/consul.${EXT}

mkdir -p "$BOOT_DESTINATION/consul"
unzip -q "$BOOT_DESTINATION"/consul.${EXT} -d "$BOOT_DESTINATION/consul"

rm -f "$BOOT_DESTINATION"/consul.${EXT}
cd "$BOOT_DESTINATION/consul"

./consul agent -dev > /dev/null &
