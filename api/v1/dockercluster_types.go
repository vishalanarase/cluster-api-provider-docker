/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

const (
	// ClusterFinalizer allows cleaning up resources associated with
	// DockerCluster before removing it from the apiserver.
	ClusterFinalizer = "dockercluster.infrastructure.cluster.x-k8s.io"
)

// DockerClusterSpec defines the desired state of DockerCluster
type DockerClusterSpec struct {
	// ControlPlaneEndpoint represents the endpoint used to communicate with the control plane.
	// +optional
	ControlPlaneEndpoint clusterv1.APIEndpoint `json:"controlPlaneEndpoint"`

	// LoadBalancerImage allows you override the load balancer image. If not specified a
	// default image will be used.
	// +optional
	LoadBalancerImage string `json:"loadbalancerImage,omitempty"`
}

// DockerClusterStatus defines the observed state of DockerCluster
type DockerClusterStatus struct {
	// Ready indicates that the cluster is ready.
	// +optional
	// +kubebuilder:default=falctrl.SetupSignalHandler()se
	Ready bool `json:"ready"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// DockerCluster is the Schema for the dockerclusters API
type DockerCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DockerClusterSpec   `json:"spec,omitempty"`
	Status DockerClusterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DockerClusterList contains a list of DockerCluster
type DockerClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DockerCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DockerCluster{}, &DockerClusterList{})
}
