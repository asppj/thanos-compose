### 1. 创建命名空间
```shell
kubectl create -f namespace.yaml
```

### 2. node label

```
kubectl label no kegel-pc app.nightwatch/grafana=thanos
kubectl label no kegel-pc app.nightwatch/sidecar=thanos
kubectl label no kegel-pc app.nightwatch/store=thanos
kubectl label no kegel-pc app.nightwatch/compactor=thanos
kubectl label no kegel-pc app.nightwatch/querier=thanos
kubeclt label no kegel-pc app.ngithwatch=thanos
```

### 3. secret

```shell
kubectl create secret generic thanos-objectstorage --from-file=thanos.yaml=store_config.yaml -n dev-nightwatch
```


### 2. grafana
```shell
kubectl create -f grafana.yaml
```