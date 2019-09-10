#!/bin/bash

if [[ -z "$BOOT_DESTINATION" ]]; then
    BOOT_DESTINATION="$HOME"
fi

EXT=tar.gz
curl -L https://archive.apache.org/dist/zookeeper/zookeeper-3.5.5/apache-zookeeper-3.5.5-bin.${EXT} -o "$BOOT_DESTINATION"/zk.${EXT}

mkdir -p "$BOOT_DESTINATION/zk"
tar xzf "$BOOT_DESTINATION"/zk.${EXT} -C "$BOOT_DESTINATION/zk" --strip-components=1

rm -f "$BOOT_DESTINATION"/zk.${EXT}
cd "$BOOT_DESTINATION/zk"

echo "tickTime=2000
dataDir=$BOOT_DESTINATION/zk/data_
clientPort=2181" > "$BOOT_DESTINATION/zk/conf/zoo.cfg"

bin/zkServer.sh start
