package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type P4 struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec P4Spec
}

type P4Spec struct {
	P4Program       string //Location of p4 program
	compilerCommand string // Compiler command
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type P4List struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Items []P4
}
