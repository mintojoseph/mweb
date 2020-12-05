# mweb

Deploy grafana with loki using terraform.

## Run terraform

``` cmdline
terraform init
terraform plan
terraform apply
```

## Configure Grafana

Get Grafana password and login using the ip from following commands.

``` cmdline
kubectl get secret --namespace grafana grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo

kubectl get services -l app.kubernetes.io/instance=grafana -n grafana -o jsonpath="{.items[0].status.loadBalancer.ingress[0].ip}
```

Get Loki ip.

``` cmdline
kubectl get svc  -l "app=loki" -o jsonpath="{.items[0].spec.clusterIP}"
```

Add above ip in the datasources like http://<< ip >>:3100.
