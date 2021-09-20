package main

import (
	"os"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	"k8s.io/client-go/kubernetes/scheme"
	log "k8s.io/klog"
)

func main() {
	obj := appsv1.StatefulSet{}
	opt := json.SerializerOptions{
		Yaml:   true,
		Pretty: true,
		Strict: true,
	}
	s := json.NewSerializerWithOptions(json.DefaultMetaFactory, scheme.Scheme, scheme.Scheme, opt)

	err := s.Encode(obj.DeepCopyObject(), os.Stdout)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
