/*
Copyright 2024.

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

package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	irsav1alpha1 "github.com/kkb0318/irsa-manager/api/v1alpha1"
	"github.com/kkb0318/irsa-manager/internal/selfhosted"
	"github.com/kkb0318/irsa-manager/internal/selfhosted/oidc"
)

// IRSASetupReconciler reconciles a IRSASetup object
type IRSASetupReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=irsa.kkb0318.github.io,resources=irsasetups,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=irsa.kkb0318.github.io,resources=irsasetups/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=irsa.kkb0318.github.io,resources=irsasetups/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the IRSASetup object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.16.3/pkg/reconcile
func (r *IRSASetupReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here

	return ctrl.Result{}, nil
}

func (r *IRSASetupReconciler) reconcile(ctx context.Context) error {
  err := reconcileSelfhosted(ctx)
	return err
}

func reconcileSelfhosted(ctx context.Context) error {
	keyPair, err := selfhosted.CreateKeyPair()
	if err != nil {
		return err
	}
	jwk, err := selfhosted.NewJWK(keyPair.PublicKey())
	if err != nil {
		return err
	}
  // get from CRs
	region := "ap-northeast-1"
	bucketName := "my-bucket-name"
	jwksFileName := "keys.json"
  var factory selfhosted.OIDCIdPFactory

	factory, err = oidc.NewAwsS3IdpFactory(
		ctx,
		region,
		bucketName,
		jwk,
		jwksFileName,
	)
	if err != nil {
		return err
	}
	issuerMeta := factory.IssuerMeta()
	discovery := factory.IdPDiscovery()
	discoveryContents := factory.IdPDiscoveryContents(issuerMeta)
	idp, err := factory.IdP(issuerMeta)
	if err != nil {
		return err
	}
	err = discovery.CreateStorage(ctx)
	if err != nil {
		return err
	}
	err = discovery.Upload(ctx, discoveryContents)
	if err != nil {
		return err
	}
	_, err = idp.Create(ctx)
	if err != nil {
		return err
	}
	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *IRSASetupReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&irsav1alpha1.IRSASetup{}).
		Complete(r)
}
