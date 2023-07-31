#!/bin/bash

function build_serverapp() {
    pushd server-app || exit 1
    docker build -t "serverapp:$1" .
    popd || exit 1
}

function build_nginx() {
    pushd nginx-image || exit 1
    docker build -t "nginx-lb:$1" .
    popd || exit 1
}

function main(){
    while [[ "$1" =~ ^-- ]]; do case $1 in
        --tag )
            shift; tag="$1"
            ;;
        --app )
            shift; app="$1"
            ;;
    esac; shift; done

    local image="$app:$tag"

    echo "Building $app with tag $tag"    
    command="build_$app $tag"
    $command
    echo "Build Finished"
}

set -e
main "$@"
