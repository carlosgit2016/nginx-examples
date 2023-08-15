#!/bin/bash

set -eo pipefail

TIME=${1:-0}
SVC=${2:-""}

function main() {
    pushd counter || exit 1
    local log_file
    log_file="$SVC.log"
    if [ -f "$log_file" ]; then
        rm $log_file
    fi
    while true; do
        response=$(curl -s -H "Host: serverapp.com" "localhost:3000/$SVC")
        jq -r '. | "Request \(.ip), \(.container_name), \(.lb_type)"'<<<$response
        jq -r '.ip'<<<"$response" >> "$log_file"
        sleep "$TIME"
    done

    popd counter || exit 1
}

main
