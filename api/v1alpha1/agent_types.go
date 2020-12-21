/*
Copyright 2020.

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

// AgentSpec defines the desired state of Agent
type AgentSpec struct {
	// Important: Run "make" to regenerate code after modifying this file

	Role string `json:"role,omitempty"`

	// Json formatted string containing the user overrides for the host's pointer ignition
	IgnitionConfigOverrides string `json:"ignition_config_overrides,omitempty" gorm:"type:text"`

	// Host installation path.
	InstallationDiskPath string `json:"installation_disk_path,omitempty"`

	// installer args
	InstallerArgs string `json:"installer_args,omitempty"`
}

// AgentStatus defines the observed state of Agent
type AgentStatus struct {
	// Important: Run "make" to regenerate code after modifying this file

	IsBootstrap bool `json:"isBootstrap"`

	// Installer version.
	InstallerVersion string `json:"installer_version,omitempty"`

	// inventory
	Inventory string `json:"inventory,omitempty" gorm:"type:text"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Agent is the Schema for the agents API
type Agent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AgentSpec   `json:"spec,omitempty"`
	Status AgentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AgentList contains a list of Agent
type AgentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Agent `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Agent{}, &AgentList{})
}
