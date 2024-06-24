// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// HTTPBootConfigSpec defines the desired state of HTTPBootConfig
type HTTPBootConfigSpec struct {
	SystemUUID        string                  `json:"systemUUID,omitempty"`
	IgnitionSecretRef *corev1.ObjectReference `json:"ignitionSecretRef,omitempty"`
	SystemIPs         []string                `json:"systemIPs,omitempty"`
	UKIURL            string                  `json:"ukiURL,omitempty"`
}

// HTTPBootConfigStatus defines the observed state of HTTPBootConfig
type HTTPBootConfigStatus struct {
	State HTTPBootConfigState `json:"state,omitempty"`
}

type HTTPBootConfigState string

const (
	HTTPBootConfigStateReady   HTTPBootConfigState = "Ready"
	HTTPBootConfigStatePending HTTPBootConfigState = "Pending"
	HTTPBootConfigStateError   HTTPBootConfigState = "Error"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="State",type=string,JSONPath=`.status.state`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
// +genclient

// HTTPBootConfig is the Schema for the httpbootconfigs API
type HTTPBootConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HTTPBootConfigSpec   `json:"spec,omitempty"`
	Status HTTPBootConfigStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// HTTPBootConfigList contains a list of HTTPBootConfig
type HTTPBootConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HTTPBootConfig `json:"items"`
}

const (
	DefaultIgnitionKey = "ignition"
	SystemUUIDIndexKey = "spec.systemUUID"
	SystemIPIndexKey   = "spec.systemIPs"
)

func init() {
	SchemeBuilder.Register(&HTTPBootConfig{}, &HTTPBootConfigList{})
}
