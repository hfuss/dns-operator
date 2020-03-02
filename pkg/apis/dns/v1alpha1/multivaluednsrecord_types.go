package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MultivalueDnsRecordSpec defines the desired state of MultivalueDnsRecord
type MultivalueDnsRecordSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// MultivalueDnsRecordStatus defines the observed state of MultivalueDnsRecord
type MultivalueDnsRecordStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MultivalueDnsRecord is the Schema for the multivaluednsrecords API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=multivaluednsrecords,scope=Namespaced
type MultivalueDnsRecord struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MultivalueDnsRecordSpec   `json:"spec,omitempty"`
	Status MultivalueDnsRecordStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MultivalueDnsRecordList contains a list of MultivalueDnsRecord
type MultivalueDnsRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MultivalueDnsRecord `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MultivalueDnsRecord{}, &MultivalueDnsRecordList{})
}
