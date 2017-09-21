package main

import (
	_ "k8s.io/kube-aggregator"
	_ "k8s.io/apiserver/pkg/storage"
	_ "k8s.io/client-go/kubernetes"
	_ "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
}
