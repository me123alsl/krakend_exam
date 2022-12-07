#!/bin/bash
OUTPUT_FILE_NAME=$1
READ_KRAKEND_JSON=$2

if [ $# -ne 2 ]; then
    echo "#######################################"
    echo "Usage: $0 <output_file_name> <krakend_json>"
    echo "#######################################"
    exit 1
fi
# READ_KRAKEND_JSON is exist
if [ ! -f "$READ_KRAKEND_JSON" ]; then
    echo "File $READ_KRAKEND_JSON does not exist"
    exit 1
fi
# OUTPUT_FILE_NAME surfix is .json
if [[ "$OUTPUT_FILE_NAME" != *.json ]]; then
    OUTPUT_FILE_NAME="$OUTPUT_FILE_NAME.json"
fi

echo "#######################################"
echo "CURRENT_DIR= $PWD"
echo "OUTPUT_JSON= $OUTPUT_FILE_NAME"
echo "KRAKEND_CONFIG= $READ_KRAKEND_JSON"
echo "#######################################"

FC_ENABLE=1 \
FC_SETTINGS="$PWD/settings" \
FC_PARTIALS="$PWD/partials" \
FC_OUT="$OUTPUT_FILE_NAME" \
krakend check -t -d -c "$PWD/$READ_KRAKEND_JSON"
