k8s-flag:
	# 打tag 哪些节点可以运行 
	kubectl tag
show-labels:
	kubectl get nodes --show-labels
PROMETHEUS_LABELS="thanos_node_type=prometheus"

add-labels:
	kubectl label nodes $(NODE) $(PROMETHEUS_LABELS)

# make NODE=demo test 传参
test:
	echo "$(NODE)"