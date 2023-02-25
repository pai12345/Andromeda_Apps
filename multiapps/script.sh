#!/bin/bash

# Build docker images
function docker_buildimage(){
    echo "==========Building Docker Images========="
    # Build docker image for app1
    cd app1
    docker build -t app1 . --compress --network=host
    cd -
    
    # Build docker image for app2
    cd app2
    docker build -t app2 . --compress --network=host
    cd -
}

# Deploy kubernetes objects
function k8s_installobjects(){
    echo "==========Deploying Kubernetes Resources========="
    cd k8s
    # deploy kubernetes namespace object
    kubectl apply -f namespace.yml
    
    # deploy kubernetes objects for app1
    kubectl apply -f app1/.
    
    # deploy kubernetes objects for app1
    kubectl apply -f app2/.
    cd -
}

# Delete all kubernetes objects
function k8s_deleteobjects(){
    echo "==========Deleting Kubernetes Resources========="
    cd k8s
    kubectl delete -f app1/.
    kubectl delete -f app2/.
    kubectl delete -f namespace.yml
    cd -
}


# Delete all docker images
function docker_deleteimages(){
    echo "==========Deleting Docker Images============="
    docker rmi $(docker images)
}

if [[ $# -eq 0 ]]; then
    echo "==================Please Select a Flag======================="
    echo "Enter -all to Install all resources"
    echo "Enter -k to Install all Kubernetes objects"
    echo "Enter -d to Build all Docker images"
    echo "Enter -un to Delete all kubernetes resources"
    exit 1
fi

if [[ $? -eq 0 ]]; then
    for KEY in "$@"; do
        case "$KEY" in
            -all)
                # "Installing all resources"
                docker_buildimage
                k8s_installobjects
            ;;
            -build)
                # "Building all Docker images"
                docker_buildimage
            ;;
            -deploy)
                # "Deploying all Kubernetes objects"
                k8s_installobjects
            ;;
            -deletebuild)
                # "Deleting all Docker images"
                docker_deleteimages
            ;;
            -deletedeploy)
                # "Deleting all Kubernetes objects"
                k8s_deleteobjects
            ;;
            *)
                echo "Invalid Flag Detected, please enter either -all, -build, -deploy, -deletebuild or -deletedeploy flags"
                exit 1
            ;;
        esac
    done
fi