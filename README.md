# `vendor-kubernetes`

This is a `go.mod` generator for consuming Kubernetes as a module.

While Kubernetes [is not meant to be consumed as a
module](https://github.com/kubernetes/kubernetes/issues/79384#issuecomment-505725449)
if you still want to do so, you'll need to pin all staging
dependencies.

## How to use it

It's required to pass a `--kubernetes-tag` argument that resembles a released
version on the Kubernetes repository.

```
~/p/g/s/g/e/vendor-kubernetes (master) > go run -mod=vendor main.go --kubernetes-tag 1.15.0 --kubernetes-path ~/projects/go/src/k8s.io
require (
  k8s.io/kubernetes v1.15.0
)

replace (
  k8s.io/api => k8s.io/api v0.0.0-20190620084959-7cf5895f2711
  k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190620085554-14e95df34f1f
  k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190612205821-1799e75a0719
  k8s.io/apiserver => k8s.io/apiserver v0.0.0-20190620085212-47dc9a115b18
  k8s.io/cli-runtime => k8s.io/cli-runtime v0.0.0-20190620085706-2090e6d8f84c
  k8s.io/client-go => k8s.io/client-go v0.0.0-20190620085101-78d2af792bab
  k8s.io/cloud-provider => k8s.io/cloud-provider v0.0.0-20190620090043-8301c0bda1f0
  k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.0.0-20190620090013-c9a0fc045dc1
  k8s.io/code-generator => k8s.io/code-generator v0.0.0-20190612205613-18da4a14b22b
  k8s.io/component-base => k8s.io/component-base v0.0.0-20190620085130-185d68e6e6ea
  k8s.io/cri-api => k8s.io/cri-api v0.0.0-20190531030430-6117653b35f1
  k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.0.0-20190620090116-299a7b270edc
  k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.0.0-20190620085325-f29e2b4a4f84
  k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.0.0-20190620085942-b7f18460b210
  k8s.io/kube-proxy => k8s.io/kube-proxy v0.0.0-20190620085809-589f994ddf7f
  k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.0.0-20190620085912-4acac5405ec6
  k8s.io/kubectl => k8s.io/kubectl v0.0.0-20190602132728-7075c07e78bf
  k8s.io/kubelet => k8s.io/kubelet v0.0.0-20190620085838-f1cb295a73c9
  k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.0.0-20190620090156-2138f2c9de18
  k8s.io/metrics => k8s.io/metrics v0.0.0-20190620085625-3b22d835f165
  k8s.io/node-api => k8s.io/node-api v0.0.0-20190620090231-de81913bee0f
  k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.0.0-20190620085408-1aef9010884e
  k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.0.0-20190620085735-97320335c8c8
  k8s.io/sample-controller => k8s.io/sample-controller v0.0.0-20190620085445-e8a5bf14f039
)
```

Optionally, you can pass a `--kubernetes-path` folder where you
have the different staging external repositories cloned, as well as
`kubernetes`. For example, my `~/projects/go/src/k8s.io` folder looks
like:

```
~ > tree -L 1 ~/projects/go/src/k8s.io
/home/ereslibre/projects/go/src/k8s.io
├── api
├── apiextensions-apiserver
├── apimachinery
├── apiserver
├── autoscaler
├── client-go
├── cli-runtime
├── cloud-provider
├── cluster-bootstrap
├── code-generator
├── component-base
├── cri-api
├── csi-translation-lib
├── kops
├── kubeadm
├── kube-aggregator
├── kube-controller-manager
├── kubectl
├── kubelet
├── kube-proxy
├── kubernetes
├── kube-scheduler
├── legacy-cloud-providers
├── metrics
├── node-api
├── sample-apiserver
├── sample-cli-plugin
├── sample-controller
├── test-infra
└── website

30 directories, 0 files
```

If you don't provide a `--kubernetes-path`, or if any staged project
is missing on your `--kubernetes-path`, `vendor-kubernetes` will clone
the repository in memory and perform the required findings. This will
be considerably slower.

As an example, let's assume that I don't have the staging project
`cluster-bootstrap` in the provided `--kubernetes-path`:

```
~/p/g/s/g/e/vendor-kubernetes (master) > rm -rf ~/projects/go/src/k8s.io/cluster-bootstrap/
~/p/g/s/g/e/vendor-kubernetes (master) > go run -mod=vendor main.go --kubernetes-tag 1.15.0 --kubernetes-path ~/projects/go/src/k8s.io
project /home/ereslibre/projects/go/src/k8s.io/cluster-bootstrap not found; cloning project https://github.com/kubernetes/cluster-bootstrap in memory
require (
  k8s.io/kubernetes v1.15.0
)

replace (
  k8s.io/api => k8s.io/api v0.0.0-20190620084959-7cf5895f2711
  k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190620085554-14e95df34f1f
  k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190612205821-1799e75a0719
  k8s.io/apiserver => k8s.io/apiserver v0.0.0-20190620085212-47dc9a115b18
  k8s.io/cli-runtime => k8s.io/cli-runtime v0.0.0-20190620085706-2090e6d8f84c
  k8s.io/client-go => k8s.io/client-go v0.0.0-20190620085101-78d2af792bab
  k8s.io/cloud-provider => k8s.io/cloud-provider v0.0.0-20190620090043-8301c0bda1f0
  k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.0.0-20190620090013-c9a0fc045dc1
  k8s.io/code-generator => k8s.io/code-generator v0.0.0-20190612205613-18da4a14b22b
  k8s.io/component-base => k8s.io/component-base v0.0.0-20190620085130-185d68e6e6ea
  k8s.io/cri-api => k8s.io/cri-api v0.0.0-20190531030430-6117653b35f1
  k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.0.0-20190620090116-299a7b270edc
  k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.0.0-20190620085325-f29e2b4a4f84
  k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.0.0-20190620085942-b7f18460b210
  k8s.io/kube-proxy => k8s.io/kube-proxy v0.0.0-20190620085809-589f994ddf7f
  k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.0.0-20190620085912-4acac5405ec6
  k8s.io/kubectl => k8s.io/kubectl v0.0.0-20190602132728-7075c07e78bf
  k8s.io/kubelet => k8s.io/kubelet v0.0.0-20190620085838-f1cb295a73c9
  k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.0.0-20190620090156-2138f2c9de18
  k8s.io/metrics => k8s.io/metrics v0.0.0-20190620085625-3b22d835f165
  k8s.io/node-api => k8s.io/node-api v0.0.0-20190620090231-de81913bee0f
  k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.0.0-20190620085408-1aef9010884e
  k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.0.0-20190620085735-97320335c8c8
  k8s.io/sample-controller => k8s.io/sample-controller v0.0.0-20190620085445-e8a5bf14f039
)
```

### `Makefile`

You can also execute `vendor-kubernetes` in the following form:

```
~/p/g/s/g/e/vendor-kubernetes (master) > KUBERNETES_TAG=1.15.0 KUBERNETES_PATH=~/projects/go/src/k8s.io make
require (
  k8s.io/kubernetes v1.15.0
)

replace (
  k8s.io/api => k8s.io/api v0.0.0-20190620084959-7cf5895f2711
  k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190620085554-14e95df34f1f
  k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190612205821-1799e75a0719
  k8s.io/apiserver => k8s.io/apiserver v0.0.0-20190620085212-47dc9a115b18
  k8s.io/cli-runtime => k8s.io/cli-runtime v0.0.0-20190620085706-2090e6d8f84c
  k8s.io/client-go => k8s.io/client-go v0.0.0-20190620085101-78d2af792bab
  k8s.io/cloud-provider => k8s.io/cloud-provider v0.0.0-20190620090043-8301c0bda1f0
  k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.0.0-20190620090013-c9a0fc045dc1
  k8s.io/code-generator => k8s.io/code-generator v0.0.0-20190612205613-18da4a14b22b
  k8s.io/component-base => k8s.io/component-base v0.0.0-20190620085130-185d68e6e6ea
  k8s.io/cri-api => k8s.io/cri-api v0.0.0-20190531030430-6117653b35f1
  k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.0.0-20190620090116-299a7b270edc
  k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.0.0-20190620085325-f29e2b4a4f84
  k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.0.0-20190620085942-b7f18460b210
  k8s.io/kube-proxy => k8s.io/kube-proxy v0.0.0-20190620085809-589f994ddf7f
  k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.0.0-20190620085912-4acac5405ec6
  k8s.io/kubectl => k8s.io/kubectl v0.0.0-20190602132728-7075c07e78bf
  k8s.io/kubelet => k8s.io/kubelet v0.0.0-20190620085838-f1cb295a73c9
  k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.0.0-20190620090156-2138f2c9de18
  k8s.io/metrics => k8s.io/metrics v0.0.0-20190620085625-3b22d835f165
  k8s.io/node-api => k8s.io/node-api v0.0.0-20190620090231-de81913bee0f
  k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.0.0-20190620085408-1aef9010884e
  k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.0.0-20190620085735-97320335c8c8
  k8s.io/sample-controller => k8s.io/sample-controller v0.0.0-20190620085445-e8a5bf14f039
)
```
