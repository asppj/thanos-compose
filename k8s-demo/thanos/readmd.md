### 1. 创建命名空间
```shell
kubectl create -f namespace.yaml
```

### 2. node label

```shell
# kubectl label no kegel-pc app.nightwatch/grafana=thanos
# kubectl label no kegel-pc app.nightwatch/sidecar=thanos
# kubectl label no kegel-pc app.nightwatch/store=thanos
# kubectl label no kegel-pc app.nightwatch/compactor=thanos
# kubectl label no kegel-pc app.nightwatch/querier=thanos
# kubeclt label no kegel-pc app.ngithwatch=thanos

# 所有节点打标签
kubectl get no|awk '{print $1}'|grep ip|xargs -t -I{} kubectl label no {} app.nightwatch=thanos
# querier
kubectl get no|awk '{print $1}'|grep ip|xargs -t -I{} kubectl label no {} app.nightwatch/querier=thanos
# sidecar
kubectl label no ip-172-32-6-118.cn-northwest-1.compute.internal app.nightwatch/sidecar=thanos
kubectl label no ip-172-32-6-135.cn-northwest-1.compute.internal app.nightwatch/sidecar=thanos
kubectl label no ip-172-32-6-109.cn-northwest-1.compute.internal app.nightwatch/sidecar=thanos

kubectl label no ip-172-32-6-118.cn-northwest-1.compute.internal app.nightwatch/querier=thanos
kubectl label no ip-172-32-6-135.cn-northwest-1.compute.internal app.nightwatch/querier=thanos
kubectl label no ip-172-32-6-109.cn-northwest-1.compute.internal app.nightwatch/querier=thanos
# compactor
kubectl label no ip-172-32-6-109.cn-northwest-1.compute.internal app.nightwatch/compactor=thanos
#grafana
kubectl label no ip-172-32-6-109.cn-northwest-1.compute.internal app.nightwatch/grafana=thanos
#store
kubectl label no ip-172-32-6-118.cn-northwest-1.compute.internal app.nightwatch/store=thanos

```

### 3. secret

```shell
kubectl create secret generic thanos-objectstorage --from-file=thanos.yaml=store_config.yaml -n dev-nightwatch
```


### 2. grafana
```shell
kubectl create -f grafana.yaml
```
### configmap
```shell
kubectl create -f rules_configmap.yaml 
kubectl create -f prometheus_configmap.yaml 
```
### rbac
```shell 
kubectl create -f rabc.yaml
```
### sidecar
```shell 
kubectl create -f sidecar.yaml
```
