# mweb

Mock HTTP service.

## Requirements

* Go
* Docker
* kubectl
* helm
* terraform
* Jenkins with following plugins- Docker Pipeline, Terraform, Kubernetes CLI. github need to be configured as 'scm'. Dockerhub credentials need to be configured as 'docker-hub-credentials'.
* minikube

### Available parameters

``` cmdline
./mweb --help
Usage of ./mweb:
  -port string
        HTTP port. HTTPPORT environment variable can also be used. (default "8080")
```

### Example

Run the application.

 ``` cmdline
 ./mweb -port 8083
```

Docker Build

``` cmdline
sudo docker build  -t  mintojoseph/mweb:1.0 .
sudo docker container run -p 8080:8080 mintojoseph/mweb:1.0
```

### Directories and files

* deployment/k8s/ - Kuberenetes yaml files for mweb application. Jenkins file is included.
* deployment/terraform/ - Terraform code for deploying mweb application using helm.
charts.
* deployment/mweb-helm/ - helm charts for mweb application.
* deployment/logging/ - Terraform code for deploying loki and grafana for monitoring using helm charts. Jenkins file is included.
* internal/ - Main program and test.
* Jenkinsfile - To build and deploy using helm charts.
* Dockerfile - To build images.
* Makefile - To build program.
