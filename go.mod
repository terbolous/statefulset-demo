module github.com/terbolous/statefulset-demo

go 1.15

require (
	k8s.io/api v0.22.2
	k8s.io/apimachinery v0.22.2
	k8s.io/client-go v0.0.0-00010101000000-000000000000
	k8s.io/klog v1.0.0
)

replace (
	k8s.io/api => k8s.io/api v0.19.10
	k8s.io/client-go => k8s.io/client-go v0.19.10
)
