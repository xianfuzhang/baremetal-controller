package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Baremetal struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BaremetalSpec   `json:"spec"`
	Status BaremetalStatus `json:"status,omitempty"`
}

// BaremetalSpec defines the desired state of Baremetal
type BaremetalSpec struct {
	// the hostname of the baremetal instance
	Hostname string `json:"hostname"`
	// id for internal use
	ID int `json:"id"`
	// the ip address of BMC
	IPMIIP string `json:"ipmiip"`
	// OS version of the baremetal instance
	OS string `json:"os"`
	// serial number of the baremetal instance
	Serial string `json:"serial"`
}

// BaremetalStatus defines the observed state of Baremetal
type BaremetalStatus struct {
	Status string `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BaremetalList contains a list of Baremetal
type BaremetalList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Baremetal `json:"items"`
}
