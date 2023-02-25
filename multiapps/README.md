# Prerequisites

- Minikube (possibly latest version)
- kubectl (possibly latest version)
- Docker (possibly latest version)

# Technical Architecture and Assumptions

## Architecture

A microservice architecture is used here with all resources deployable in a Minikube kuberentes cluster.

I have chosen service to service communication mechanism for interpod communication within a defined kuberentes namespace.

- I have chosen this interpod communication approach as it is supported by default in every Minikube cluster without requiring the need to enable any additional plugins or DNS configurations as is in the case of using an ingress controller.
- Both the application services are of type `NodePort` and are mapped to a port on the host.
- This approach also makes it easier to test both the applications without any hinderance of extra infrastructure configurations.

The first application is referred in the project as `app1` and the second application is referred as `app2`.

## First application - app1

- The first application `app1` is responsible for accepting dynamic messages from the End User and displays a JSON document as response eg: `{ "id": "1", "message": "Hello world" }` when visited with a HTTP client.

- app1 is developed using python based FASTAPI framework which is a lightning fast, high performance and production ready framework served through Uvicorn, an ASGI web server implementation for python.

- The application has been implemented with both `GET` and `POST` REST based capabilities to fetch and feed the message.

- The application takes care of persisting the message by writing the message contents to a `cache.txt` file as long as the application is healthy. Reload of the application will delete the original message as the cache.txt file is specific to app1 itself. The txt file will contain the message "Hello world" by default but can be changed to a different message with a POST request as explained below.

- app1 makes use of aiofiles python library for all asynchronous file operations.

### GET Request

On first visit of app1 using GET request, the application reads the contents of `cache.txt` file and will display `{ "id": "1", "message": "Hello world" }`. If the contents of the cache.txt file has been changed eg. after a POST request, then the message displayed will also change.

```
eg: Display JSON Document

curl -i -X GET 'http://localhost:30007'

HTTP/1.1 200 OK
date: Sat, 11 Feb 2023 23:02:04 GMT
server: uvicorn
content-length: 34
content-type: application/json

{"id":"1","message":"Hello World"}
```

### POST request

Using the POST request, the End User can pass a dynamic message.

```
eg: Changing message to 'Cloud Native Computing Foundation'

curl -i -X POST -H "Content-Type:application/json" -d '{"id":"1","message":"Cloud Native Computing Foundation"}' 'http://localhost:30007'

HTTP/1.1 200 OK
date: Sat, 11 Feb 2023 23:05:23 GMT
server: uvicorn
content-length: 56
content-type: application/json

{"id":"1","message":"Cloud Native Computing Foundation"}

```

If we use the GET request again, the new message will be displayed in the JSON Document.

```
eg: Display new message

curl -i -X GET 'http://localhost:30007'

HTTP/1.1 200 OK
date: Sat, 11 Feb 2023 23:06:38 GMT
server: uvicorn
content-length: 56
content-type: application/json

{"id":"1","message":"Cloud Native Computing Foundation"}

```

## Second Application - app2

- The second application `app2`, is responsible for utilizing the first application and displaying its fully reversed message text rendered dynamically. Hence app2 has been implemented for only `GET` request.

- Similar to app1, app2 is developed using FASTAPI framework served with Uvicorn.

- app2 uses aiohttp, an asynchronous http client to retrive the JSON document from app1.

- It then reverses the message and displays the JSON document containing the reversed message to the End User.

- There is no data persistence for app2 as its sole responsibilty is to fetch message from app1 and display the fully reversed message to the End User.

```
eg: Display reversed message

curl -i -X GET 'http://localhost:30008'

HTTP/1.1 200 OK
date: Sat, 11 Feb 2023 23:10:09 GMT
server: uvicorn
content-length: 34
content-type: application/json

{"id":"1","message":"dlroW olleH"}

```

## End to End Working

- The End User will first submit a `POST` request to app1 with the desired message.
- This message is persisted in app1. The End User can send as many POST requests with desired messages and the last message will be persisted by app1.
- The End User can also access app1 via GET request and will display the current message persisted by app1.

- Then the enduser will access the second application `app2` via an HTTP GET request to view the reversed message.
- app2 sends a GET request to app1 for fetching the message.
- app1 sends back the JSON document with the message as response to app2.
- Once app2 recieves the JSON document having the message, it then reverses the message and displays the JSON document having the reveresed text message.

# Installation

The entire project is equipped with an automation script that automates the build of all container images and deploys all required kuberentes objects.

In a nut shell the script works as below.

- Build all container images and Deploy all kubernetes objects: `./script.sh -all`

- Build all container images: `./script.sh -build`

- Deploy all Kubernetes objects: `./script.sh -deploy`

- Delete all container images: `./script.sh -deletebuild`

- Delete all kubernetes objects: `./script.sh -deletedeploy`

A detailed information can be found in the below sections.

Both `app1` and `app2` application folders contain a Dockerfile that can be used to create container images. These container images need to be built and available locally inside the Minikube environment where the kuberentes objects are deployed.

The entire project is shipped as a compressed archive `project.tar.gz` as required by the challenge. This archive will need to be coppied and extracted to the Minikube environment for accessing all files.

## Auto Build Container Images

To build all container images run `./script.sh -build` in the terminal.

The `-build` flag instructs the automation script to build the image for both app1 and app2 tagged to latest version.

Once successfull the container images can be verified from the termianl using `docker images` command.

```
eg: docker images

REPOSITORY   TAG       IMAGE ID       CREATED          SIZE
app2         latest    e914c536d1da   6 seconds ago    95.1MB
app1         latest    c5c9e2d6fea2   15 seconds ago   85.3MB
```

## Auto Deploy Kubernetes resources

To deploy all kubernetes resources run `./script.sh -deploy` in the terminal.

The `-deploy` flag instructs the automation script to first deploy the kuberentes namespace object followed by all the kubernetes objects for app1 and app2 resources.

Once completed, the kubernetes objects deployed by the script can be verified from the terminal using using the `kubectl get all -n dev` command.

```
eg: kubectl get all -n dev

pod/deploy-app1-6b996ff4c7-nx5kq   1/1     Running   0          27s
pod/deploy-app2-6f99f45bbc-d6lms   1/1     Running   0          21s

NAME               TYPE       CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
service/svc-app1   NodePort   10.100.4.109     <none>        8080:30007/TCP   27s
service/svc-app2   NodePort   10.100.190.217   <none>        8081:30008/TCP   21s

NAME                          READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/deploy-app1   1/1     1            1           29s
deployment.apps/deploy-app2   1/1     1            1           23s

NAME                                     DESIRED   CURRENT   READY   AGE
replicaset.apps/deploy-app1-6b996ff4c7   1         1         1       29s
replicaset.apps/deploy-app2-6f99f45bbc   1         1         1       23s
```

## Auto Delete Kubernetes Objects and Container Images

To delete all kubernetes objects deployed by the automation script, run `./script.sh -deletedeploy` command in the terminal. The `-deletedeploy` flag instructs the script to delete all kubernetes objects for both app1 and app2.

To delete all container images built by the automation script, run `./script.sh -deletebuild` command in the terminal. The `-deletebuild` flag instructs the script to delete all the container images for app1 and app2.

# Areas of future improvement

- `Kuberentes Health checks - Liveness, Readiness and Startup Probes`: I have currently not included this in the yaml files as I require the information on CPU and Memory of the worker node.`
- `Kubernetes QOS(Quality of Service)`: Used for resource requests and limits, which allow kuberentes to control CPU and memory resource constraints. Similar to the above, I will require information on CPU and Memory of the worker node.
- `Ingress Controller`: Used for managing external access (Layer 7) to the services in the cluster. This will require dns configuration and additional plugins to be enabled in Minikube.
- `Pod Disruption Budget(PDB)`: Used for controlling voluntary disruptions for pod replicas. In the solution I have enabled only 1 pod replica and did not want to complicate any tests.
- `Horizontal Pod Autoscaler (HPA)`: Used for horizontally autoscalling pods based on CPU and memory utlisation. I will require CPU and memory information of the worker node for deploying this resource.
- `Network Policy`: Used for controlling network traffic and network isolations for the pods. This will require network plugins eg. Calico to be enabled in Minikube.

# Conclusion

The solution implemented aims on achieving the "Containerized microservice challenge" and focuses on simplicity, transparency, core kuberentes concepts and without the need of any extra infrastructure configurations or setups. This solution also demonstrates the core aspects of Kubernetes service to service interpod communication at Layer 4 without the need to configure any extra DNS records. All applications for the challenge are developed using latest Cutting Edge python frameworks that focus on asynchronous programming models with efficient CPU and memory utilisations. Overall I personally really enjoyed the challenge and would love to contribute more to your team and organisation.
