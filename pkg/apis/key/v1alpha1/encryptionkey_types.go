package v1alpha1

import (
	"github.com/levigross/encrypted-secret-operator/pkg/common/status"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EncryptionKeySpec defines the desired state of EncryptionKey
type EncryptionKeySpec struct {

	// The type of encryption key that you used to encrypt the data
	// +kubebuilder:validation:Enum=GPG;SSH
	Type string `json:"type"`

	PrivateEncryptionKey []byte `json:"privateEncryptionKey"`

	// PublicEncryptionKey should be used if you want to validate that the private key
	// matches the public key
	PublicEncryptionKey []byte `json:"privateEncryptionKey,omitempty"`

	// If you supply an encrypted key, we can decrypt it when using it.
	// *Note* we may no effort to scrub this value from memory afterwards.
	Passphrase string `json:"passphrase,omitempty"`
}

// EncryptionKeyStatus defines the observed state of EncryptionKey
type EncryptionKeyStatus struct {
	KeyStatus []status.KeyStatus `json:"keyStatus"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EncryptionKey is the Schema for the encryptionkeys API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=encryptionkeys,scope=Namespaced
type EncryptionKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EncryptionKeySpec   `json:"spec,omitempty"`
	Status EncryptionKeyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EncryptionKeyList contains a list of EncryptionKey
type EncryptionKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EncryptionKey `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EncryptionKey{}, &EncryptionKeyList{})
}
