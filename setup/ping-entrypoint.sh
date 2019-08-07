#!/bin/bash
[[ -n $DEBUG ]] && set -x -e

cwd=`dirname "$0"`
expr "$0" : "/.*" > /dev/null || cwd=`(cd "$cwd" && pwd)`
UPLOAD_FILE="$cwd/ping"

function setup_release_base(){
  account=$1
  loc=$2 
  echo "Creating an Azure storage account: $account (region: $loc)"
  az storage account create --name $account \
    --location $loc \
    --resource-group $RESOURCE_GROUP \
    --sku Standard_LRS \
    --kind StorageV2

  # Get Storage Key
  access_key=$(az storage account keys list \
    --account-name $account \
    --resource-group $RESOURCE_GROUP \
    --output tsv |head -1 | awk '{print $3}')

  echo "Create a container: \$root"
  az storage container create  \
    --name "\$root" \
    --account-name $account \
    --account-key $access_key \
    --public-access container
}


function setup_endpoint(){
  account=$1
  loc=$2 
  echo "Creating an Azure storage account: $account (region: $loc)"
  az storage account create --name $account \
    --location $loc \
    --resource-group $RESOURCE_GROUP \
    --sku Standard_LRS \
    --kind StorageV2

  # Get Storage Key
  access_key=$(az storage account keys list \
    --account-name $account \
    --resource-group $RESOURCE_GROUP \
    --output tsv |head -1 | awk '{print $3}')

  echo "Create a container: \$root"
  az storage container create  \
    --name "\$root" \
    --account-name $account \
    --account-key $access_key \
    --public-access container

  echo "Uploading a file ($UPLOAD_FILE) ping to \$root"
  az storage blob upload \
    --account-name $account \
    --account-key $access_key \
    --container-name "\$root" \
    --name ping \
    --file $UPLOAD_FILE
}

function create_resource_group(){
  echo "Creating Resource Group: $RESOURCE_GROUP"
  az group create --name $RESOURCE_GROUP --location "japaneast"
}

#
# Main
#

cat << EOD | tee
EXPLANATION
...
EOD

create_resource_group

for r in $REGION_LIST
do
  echo "Target region: $r **************"
  storageaccount="azping$r"
  setup_endpoint $storageaccount $r
done

for r in $REGION_LIST
do
  url="http://azping$r.blob.core.windows.net/ping"
  echo "Access test: $url"
  curl -s $url
done

setup_release_base $RELEASE_ACCOUNT "japaneast"

echo "Done"
