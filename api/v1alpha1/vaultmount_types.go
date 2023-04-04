// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// VaultMountSpec defines the desired state of VaultMountSpec
type VaultMountSpec struct {
	// VaultAuthRef of the VaultAuth resource
	// If no value is specified the Operator will default to the `default` VaultAuth,
	// configured in its own Kubernetes namespace.
	VaultAuthRef string `json:"vaultAuthRef,omitempty"`

	// Namespace to auth to in Vault
	Namespace string `json:"namespace,omitempty"`

	// Mount to use when authenticating to auth method.
	Path string `json:"path"`

	// Type must be a host string,
	// +kubebuilder:validation:Required
	Type string `json:"type,omitempty"`

	// Type must be a host string,
	// +kubebuilder:validation:Optional
	Description string `json:"description,omitempty"`

}

// VaultMountStatus defines the observed state of VaultMountSpec
type VaultMountStatus struct {
	// Valid auth mechanism.
	Valid bool   `json:"valid"`
	Error string `json:"error"`
	Path  string `json:"path"`
}

// +kubebuilder:object:root=true
//+kubebuilder:subresource:status

// VaultMount is the Schema for the VaultMounts API
type VaultMount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VaultMountSpec   `json:"spec,omitempty"`
	Status VaultMountStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// VaultMountList contains a list of VaultMount
type VaultMountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VaultMount `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VaultMount{}, &VaultMountList{})
}
