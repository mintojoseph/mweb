# mweb

Mock HTTP service.

## Requirements

Go
Docker
kubectl
helm
terraform

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

### Directories

k8s - Kuberenetes yaml files.
terraform - Terraform code for deploying application using helm charts.
logging - Terraform code for deploying loki and grafana for monitoring using helm charts.
