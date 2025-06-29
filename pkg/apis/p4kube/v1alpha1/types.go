package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type P4 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   P4Spec   `json:"spec"`
	Status P4Status `json:"status"`
}

type P4Spec struct {
	P4Program         string `json:"p4ProgramLocation"`
	CompilerDirectory string `json:"compilerDirectory"`
	CompilerCommand   string `json:"compilerCommand"`
	TargetNode        string `json:"targetNode"`
	NetworkFunction   string `json:"networkFunction"`
	DeploymentPhase   string `json:"deploymentPhase"`
}

type P4Status struct {
	Progress        string `json:"progress"`
	Node            string `json:"node"`
	DeploymentPhase string `json:"deploymentPhase"`
	NetworkFunction string `json:"networkFunction"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type P4List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []P4 `json:"items"`
}
