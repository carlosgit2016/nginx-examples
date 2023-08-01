#!/bin/bash

function build_serverapp() {
    pushd server-app || exit 1
    docker build -t "serverapp:$1" .
    popd || exit 1
}

function build_nginx() {
    pushd nginx-image || exit 1
    local tag
    docker build -t "nginx-lb:$tag" .
    popd || exit 1
}

function main(){
    while [[ "$1" =~ ^-- ]]; do case $1 in
        --tag )
            shift; tag="$1";
            ;;
        --app )
            shift; app="$1"
            ;;
    esac; shift; done

    if [ -z "$tag" ]; then
        tag="latest"
    fi

    if [ -z "$app" ]; then
        echo "App required" && exit 1
    fi

    local image="$app:$tag"

    echo "Building $app with tag $tag"    
    command="build_$app $tag"
    $command
    echo "Build Finished"
}

set -e
main "$@"
