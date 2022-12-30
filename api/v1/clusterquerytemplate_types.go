/*
Copyright 2022.

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
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ClusterQueryTemplateSpec defines the desired state of ClusterQueryTemplate
type ClusterQueryTemplateSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Providers Providers `json:"providers,omitempty" protobuf:"bytes,1,opt,name=providers"`
}

type Providers struct {
	Name       string       `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	Prometheus []Prometheus `json:"prometheus,omitempty" protobuf:"bytes,2,opt,name=prometheus"`
	Wavefront  []Wavefront  `json:"wavefront,omitempty" protobuf:"bytes,3,opt,name=wavefront"`
}

type Prometheus struct {
	Name       string               `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	Queries    []QueryTemplateValue `json:"queries,omitempty" protobuf:"bytes,2,opt,name=queries"`
	TimeLength string               `json:"timeLength,omitempty" protobuf:"bytes,3,opt,name=timeLength"`
	Step       string               `json:"step,omitempty" protobuf:"bytes,4,opt,name=step"`
	Address    string               `json:"address,omitempty" protobuf:"bytes,5,opt,name=address"`
}

type Wavefront struct {
	Name    string               `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	Queries []QueryTemplateValue `json:"queries,omitempty" protobuf:"bytes,2,opt,name=queries"`
	Address string               `json:"address,omitempty" protobuf:"bytes,4,opt,name=address"`
}

type QueryTemplateValue struct {
	Name  string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	Query string `json:"query,omitempty" protobuf:"bytes,2,opt,name=query"`
}

// ClusterQueryTemplateStatus defines the observed state of ClusterQueryTemplate
type ClusterQueryTemplateStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:scope=Cluster

// ClusterQueryTemplate is the Schema for the clusterquerytemplates API
type ClusterQueryTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterQueryTemplateSpec   `json:"spec,omitempty"`
	Status ClusterQueryTemplateStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ClusterQueryTemplateList contains a list of ClusterQueryTemplate
type ClusterQueryTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterQueryTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterQueryTemplate{}, &ClusterQueryTemplateList{})
}
