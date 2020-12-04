# mweb

Mock HTTP service.

## Requirements

Go

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

Deploy kubernetes service

``` cmdline
#   kubectl apply -f deployment.yml
#   kubectl apply -f metallb.yml
#   kubectl apply -f service.yml
```
