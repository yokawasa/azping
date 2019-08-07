az account list-locations -o tsv |awk -F\t '{print $5}'
