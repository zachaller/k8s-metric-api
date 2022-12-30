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
	"k8s.io/apimachinery/pkg/runtime/schema"
)

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// QueryTemplate is the Schema for the querytemplates API
type QueryTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterQueryTemplateSpec   `json:"spec,omitempty"`
	Status ClusterQueryTemplateStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// QueryTemplateList contains a list of QueryTemplate
type QueryTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []QueryTemplate `json:"items"`
}

func (qt *QueryTemplate) GetGroupVersionResource() schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    "query.metrics-api.io",
		Version:  "v1",
		Resource: "querytemplates",
	}
}

func init() {
	SchemeBuilder.Register(&QueryTemplate{}, &QueryTemplateList{})
}
