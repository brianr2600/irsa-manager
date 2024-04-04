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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// IRSASetupKind represents the kind attribute of an IRSASetup resource.
	IRSASetupKind = "IRSASetup"
)

// IRSASetupSpec defines the desired state of IRSASetup
type IRSASetupSpec struct {
	// Mode specifies the mode of operation. Can be either "selfhosted" or "eks".
	Mode string `json:"mode"`

	// Discovery configures the IdP Discovery process, essential for setting up IRSA by locating
	// the OIDC provider information.
	Discovery Discovery `json:"discovery"`

	// Auth contains authentication configuration details.
	Auth Auth `json:"auth,omitempty"`
}

// Discovery holds the configuration for IdP Discovery, which is crucial for locating
// the OIDC provider in a self-hosted environment.
type Discovery struct {
	// S3 specifies the AWS S3 bucket details where the OIDC provider's discovery information is hosted.
	S3 S3Discovery `json:"s3,omitempty"`
}

// S3Discovery contains the specifics of the S3 bucket used for hosting OIDC provider discovery information.
type S3Discovery struct {
	// Region denotes the AWS region where the S3 bucket is located.
	Region string `json:"region"`

	// BucketName is the name of the S3 bucket that hosts the OIDC discovery information.
	BucketName string `json:"bucketName"`
}

// Auth holds the authentication configuration details.
type Auth struct {
	// SecretRef specifies the reference to the Kubernetes secret containing authentication details.
	SecretRef SecretRef `json:"secretRef"`
}

// SecretRef contains the reference to a Kubernetes secret.
type SecretRef struct {
	// Name specifies the name of the secret.
	Name string `json:"name"`

	// Namespace specifies the namespace of the secret.
	Namespace string `json:"namespace,omitempty"`
}

// IRSASetupStatus defines the observed state of IRSASetup
type IRSASetupStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// IRSASetup represents a configuration for setting up IAM Roles for Service Accounts (IRSA) in a Kubernetes cluster.
type IRSASetup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IRSASetupSpec   `json:"spec,omitempty"`
	Status IRSASetupStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// IRSASetupList contains a list of IRSASetup
type IRSASetupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IRSASetup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&IRSASetup{}, &IRSASetupList{})
}