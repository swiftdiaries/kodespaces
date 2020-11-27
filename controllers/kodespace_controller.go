/*


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

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	kodev1alpha1 "github.com/swiftdiaries/kodespaces/api/v1alpha1"
)

// KodespaceReconciler reconciles a Kodespace object
type KodespaceReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=kode.swiftdiaries.com,resources=kodespaces,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=kode.swiftdiaries.com,resources=kodespaces/status,verbs=get;update;patch

func (r *KodespaceReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("kodespace", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *KodespaceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&kodev1alpha1.Kodespace{}).
		Complete(r)
}
