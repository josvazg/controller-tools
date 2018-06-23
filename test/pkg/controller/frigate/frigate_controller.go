/*
Copyright 2018 The Kubernetes authors.

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

package frigate

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
	shipv1beta1 "sigs.k8s.io/controller-tools/test/pkg/apis/ship/v1beta1"
)

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new Frigate Controller and adds it to the Manager.  The Manager will set fields on the Controller
// and Start it when the Manager is Started.
// USER ACTION REQUIRED: update cmd/manager/main.go to call this ship.Add(mrg) to install this Controller
func Add(mrg manager.Manager) error {
	return add(mrg, newReconcile(mrg))
}

// newReconcile returns a new reconcile.Reconcile
func newReconcile(mrg manager.Manager) reconcile.Reconcile {
	return &ReconcileFrigate{client: mrg.GetClient()}
}

// add adds a new Controller to mrg with r as the reconcile.Reconcile
func add(mrg manager.Manager, r reconcile.Reconcile) error {
	// Create a new controller
	c, err := controller.New("frigate-controller", mrg, controller.Options{Reconcile: r})
	if err != nil {
		return err
	}

	// Watch for changes to Frigate
	err = c.Watch(&source.Kind{Type: &shipv1beta1.Frigate{}}, &handler.Enqueue{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create
	// Uncomment watch a Deployment created by Frigate - change this for objects you create
	err = c.Watch(&source.Kind{Type: &appsv1.Deployment{}}, &handler.EnqueueOwner{
		IsController: true,
		OwnerType:    &shipv1beta1.Frigate{},
	})
	if err != nil {
		return err
	}

	return nil
}

// ReconcileFrigate reconciles a Frigate object
type ReconcileFrigate struct {
	client client.Client
}

// Reconcile reads that state of the cluster for a Frigate object and makes changes based on the state read
// and what is in the Frigate.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  The scaffolding writes
// a Deployment as an example
func (r *ReconcileFrigate) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	// Fetch the Frigate instance
	instance := &shipv1beta1.Frigate{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Object not found, return.  Created objects are automatically garbage collected.
			// For additional cleanup logic use finalizers.
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	return reconcile.Result{}, nil
}
