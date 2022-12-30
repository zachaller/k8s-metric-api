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

package controllers

import (
	"context"
	prometheusv1 "github.com/zachaller/k8s-metrics-api/apiserver/pkg/apis/prometheus/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	queryv1 "github.com/zachaller/k8s-metrics-api/api/v1"
)

// QueryReconciler reconciles a Query object
type QueryReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=query.metrics-api.io,resources=queries,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=query.metrics-api.io,resources=queries/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=query.metrics-api.io,resources=queries/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Query object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.1/pkg/reconcile
func (r *QueryReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	var metricQuery queryv1.Query
	err := r.Client.Get(ctx, req.NamespacedName, &metricQuery)
	if err != nil {
		return ctrl.Result{}, err
	}

	var metricQueryRun prometheusv1.MetricQueryRun
	r.Client.Get(ctx, req.NamespacedName, &metricQueryRun) // ignore error because we want to create it if it doesn't exist
	if metricQueryRun.Name == metricQuery.Name {
		return ctrl.Result{}, err
	}

	annotations := metricQuery.GetAnnotations()
	delete(annotations, "kubectl.kubernetes.io/last-applied-configuration")

	metricQueryRun = prometheusv1.MetricQueryRun{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:        metricQuery.Name,
			Namespace:   metricQuery.Namespace,
			Labels:      metricQuery.Labels,
			Annotations: annotations,
			OwnerReferences: []metav1.OwnerReference{{
				APIVersion: metricQuery.APIVersion,
				Kind:       metricQuery.Kind,
				Name:       metricQuery.Name,
				UID:        metricQuery.UID,
			}},
		},
		Spec: prometheusv1.MetricQueryRunSpec{},
	}

	//client, _, err := kubeclient.NewKubeClient()
	//dynClient.Resource(metricQueryRun.GetGroupVersionResource()).Namespace(metricQuery.Namespace).Create(ctx, &metricQueryRun., metav1.CreateOptions{})
	//r.Client.Create(ctx, &metricQueryRun)
	err = r.Client.Create(ctx, &metricQueryRun)
	if err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *QueryReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&queryv1.Query{}).
		Complete(r)
}
