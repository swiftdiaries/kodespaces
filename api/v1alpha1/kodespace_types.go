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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// KodespaceSpec defines the desired state of Kodespace
type KodespaceSpec struct {
	// IDESpec for codespace - options are Jupyter, Theia IDE
	Ide IDESpec `json:"ide,omitempty"`
}

// KodespaceStatus defines the observed state of Kodespace
type KodespaceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// Kodespace is the Schema for the kodespaces API
type Kodespace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KodespaceSpec   `json:"spec,omitempty"`
	Status KodespaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KodespaceList contains a list of Kodespace
type KodespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Kodespace `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Kodespace{}, &KodespaceList{})
}
