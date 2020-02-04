package v1alpha1

import (
	"github.com/levigross/encrypted-secret-operator/pkg/common/status"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EncryptedSecretSpec defines the desired state of EncryptedSecret
type EncryptedSecretSpec struct {
	corev1.Secret `json:",inline"`
}

// EncryptedSecretStatus defines the observed state of EncryptedSecret
type EncryptedSecretStatus struct {
	SecretStatus []status.KeyStatus `json:"secretStatus"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EncryptedSecret is the Schema for the encryptedsecrets API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=encryptedsecrets,scope=Namespaced
type EncryptedSecret struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EncryptedSecretSpec   `json:"spec,omitempty"`
	Status EncryptedSecretStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EncryptedSecretList contains a list of EncryptedSecret
type EncryptedSecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EncryptedSecret `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EncryptedSecret{}, &EncryptedSecretList{})
}
