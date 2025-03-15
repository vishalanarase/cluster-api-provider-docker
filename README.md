# cluster-api-provider-docker


❯ kubebuilder init --domain  cluster.x-k8s.io --repo github.com/vishalanarase/cluster-api-provider-docker
INFO Writing kustomize manifests for you to edit...
INFO Writing scaffold for you to edit...
INFO Get controller runtime:
$ go get sigs.k8s.io/controller-runtime@v0.17.3
INFO Update dependencies:
$ go mod tidy
Next: define a resource with:
$ kubebuilder create api


❯ kubebuilder create api --group infrastructure --version v1 --kind DockerCluster
INFO Create Resource [y/n]
y
INFO Create Controller [y/n]
y
INFO Writing kustomize manifests for you to edit...
INFO Writing scaffold for you to edit...
INFO api/v1/dockercluster_types.go
INFO api/v1/groupversion_info.go
INFO internal/controller/suite_test.go
INFO internal/controller/dockercluster_controller.go
INFO internal/controller/dockercluster_controller_test.go
INFO Update dependencies:
$ go mod tidy
INFO Running make:
$ make generate
mkdir -p /Users/vishal/workspace/vishalanarase/cluster-api-provider-docker/bin
Downloading sigs.k8s.io/controller-tools/cmd/controller-gen@v0.14.0
go: downloading sigs.k8s.io/controller-tools v0.14.0
go: downloading github.com/spf13/cobra v1.8.0
go: downloading github.com/fatih/color v1.16.0
go: downloading k8s.io/api v0.29.0
go: downloading k8s.io/apimachinery v0.29.0
go: downloading k8s.io/apiextensions-apiserver v0.29.0
go: downloading github.com/gobuffalo/flect v1.0.2
go: downloading github.com/mattn/go-colorable v0.1.13
go: downloading github.com/mattn/go-isatty v0.0.20
go: downloading golang.org/x/mod v0.14.0
go: downloading github.com/go-logr/logr v1.3.0
go: downloading golang.org/x/sys v0.15.0
/Users/vishal/workspace/vishalanarase/cluster-api-provider-docker/bin/controller-gen-v0.14.0 object:headerFile="hack/boilerplate.go.txt" paths="./..."
Next: implement your new API and generate the manifests (e.g. CRDs,CRs) with:
$ make manifests


❯ kubebuilder create api --group infrastructure --version v1 --kind DockerMachine
INFO Create Resource [y/n]
y
INFO Create Controller [y/n]
y
INFO Writing kustomize manifests for you to edit...
INFO Writing scaffold for you to edit...
INFO api/v1/dockermachine_types.go
INFO api/v1/groupversion_info.go
INFO internal/controller/suite_test.go
INFO internal/controller/dockermachine_controller.go
INFO internal/controller/dockermachine_controller_test.go
INFO Update dependencies:
$ go mod tidy
INFO Running make:
$ make generate
/Users/vishal/workspace/vishalanarase/cluster-api-provider-docker/bin/controller-gen-v0.14.0 object:headerFile="hack/boilerplate.go.txt" paths="./..."
Next: implement your new API and generate the manifests (e.g. CRDs,CRs) with:
$ make manifests


