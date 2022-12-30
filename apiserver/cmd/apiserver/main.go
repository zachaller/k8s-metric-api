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

package main

import (
	"k8s.io/klog"
	"sigs.k8s.io/apiserver-runtime/pkg/builder"

	//"sigs.k8s.io/apiserver-runtime/pkg/experimental/storage/filepath"
	memoryStorage "github.com/zachaller/k8s-metrics-api/apiserver/pkg/storage"

	// +kubebuilder:scaffold:resource-imports
	prometheusv1 "github.com/zachaller/k8s-metrics-api/apiserver/pkg/apis/prometheus/v1"
)

func main() {
	cmd, err := builder.APIServer.
		// +kubebuilder:scaffold:resource-register
		WithOpenAPIDefinitions("metrics", "v0.0.0", prometheusv1.GetOpenAPIDefinitions).
		//WithResourceAndHandler(&prometheusv1.MetricQueryRun{}, filepath.NewJSONFilepathStorageProvider(&prometheusv1.MetricQueryRun{}, "data")).
		WithResourceAndHandler(&prometheusv1.MetricQueryRun{}, memoryStorage.NewMemoryStorageProvider(&prometheusv1.MetricQueryRun{})).
		WithLocalDebugExtension().
		Build()
	cmd.Execute()
	if err != nil {
		klog.Fatal(err)
	}
}
