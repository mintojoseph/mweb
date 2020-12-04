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

 ``` cmdline
 ./mweb -port 8083
```

``` cmdline
sudo docker build  -t  mintojoseph/mweb:1.0 .
sudo docker container run -p 8080:8080 mintojoseph/mweb:1.0
```
