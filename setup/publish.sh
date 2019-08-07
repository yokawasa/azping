#!/bin/bash
[[ -n $DEBUG ]] && set -x -e

cwd=`dirname "$0"`
expr "$0" : "/.*" > /dev/null || cwd=`(cd "$cwd" && pwd)`
source $cwd/env.sh

release_account="azpingrelease"

#
# Main
#
files="azping_linux_amd64 azping_darwin_amd64 azping_windows_amd64"

# Get Storage Key
access_key=$(az storage account keys list \
    --account-name $release_account \
    --resource-group $RESOURCE_GROUP \
    --output tsv |head -1 | awk '{print $3}')

for f in $files
do 
  upload_file="$cwd/../bin/$f"
  echo "Uploading a file ($upload_file) ping to \$root"
  az storage blob upload \
    --account-name $release_account \
    --account-key $access_key \
    --container-name "\$root" \
    --name $f \
    --file $upload_file
done
