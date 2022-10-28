/*
Copyright 2022 The members of the #B4mad.Net.

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

// SPDX-License-Identifier: GPL-3.0-or-later

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PitmasterSpec defines the desired state of Pitmaster
type PitmasterSpec struct {
	// CrewchiefUsernames is a list of Crewchief usernames that want to use Pitcrew
	CrewchiefUsernames []string `json:"crewchiefusernames,omitempty"`
}

// PitmasterStatus defines the observed state of Pitmaster
type PitmasterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Pitmaster is the Schema for the pitmasters API
type Pitmaster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PitmasterSpec   `json:"spec,omitempty"`
	Status PitmasterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PitmasterList contains a list of Pitmaster
type PitmasterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Pitmaster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Pitmaster{}, &PitmasterList{})
}
