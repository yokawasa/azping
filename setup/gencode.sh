#!/bin/bash
[[ -n $DEBUG ]] && set -x -e

for r in $REGION_LIST
do
  host="azping$r.blob.core.windows.net"
  echo "\"$r\":          \"$host\","
done

echo "Done"
