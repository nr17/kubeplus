/*
Copyright 2017 The Kubernetes Authors.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Foo is a specification for a Foo resource
type Postgres struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PostgresSpec   `json:"spec"`
	Status PostgresStatus `json:"status"`
}

// PostgresSpec is the spec for a Foo resource
type PostgresSpec struct {
	DeploymentName string `json:"deploymentName"`
	Image string `json:"image"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	Replicas       *int32 `json:"replicas"`
	Commands []string `json:"commands"`
}

// FooStatus is the status for a Foo resource
type PostgresStatus struct {
	AvailableReplicas int32 `json:"availableReplicas"`
	ActionHistory []string `json:"actionHistory"`
	VerifyCmd string `json:"verifyCommand"`
	ServiceIP string `json:"serviceIP"`
	ServicePort string `json:"servicePort"`
	Status string `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FooList is a list of Foo resources
type PostgresList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Postgres `json:"items"`
}
