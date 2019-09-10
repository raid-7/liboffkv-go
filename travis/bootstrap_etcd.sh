#!/bin/bash

if [[ -z "$BOOT_DESTINATION" ]]; then
    BOOT_DESTINATION="$HOME"
fi

TRIPLET="$1"-amd64
EXT=$2
curl -L https://github.com/etcd-io/etcd/releases/download/v3.3.13/etcd-v3.3.13-${TRIPLET}.${EXT} -o "$BOOT_DESTINATION"/etcd.${EXT}

mkdir -p "$BOOT_DESTINATION/etcd"
if [[ "$EXT" == "tar.gz" ]]; then
	tar xzf "$BOOT_DESTINATION"/etcd.${EXT} -C "$BOOT_DESTINATION/etcd" --strip-components=1
elif [[ "$EXT" == "zip" ]]; then
	unzip "$BOOT_DESTINATION"/etcd.${EXT} -d "$BOOT_DESTINATION/etcd"
	f=("$BOOT_DESTINATION/etcd"/*) && mv "$BOOT_DESTINATION/etcd"/*/* "$BOOT_DESTINATION/etcd" && rmdir "${f[@]}"
else
	exit 1
fi

rm -f "$BOOT_DESTINATION"/etcd.${EXT}
cd "$BOOT_DESTINATION/etcd"

mkdir _data
./etcd --data-dir "$(pwd)/_data" > /dev/null &
