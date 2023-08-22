package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var SchemeGroupVersion = schema.GroupVersion{
	Group:   "p4kube.comnets",
	Version: "v1alpha1",
}

var (
	SchemeBuilder runtime.SchemeBuilder
	AddToScheme   = SchemeBuilder.AddToScheme
)

func init() {
	// This func is called only once as soon as the package (v1alpha1) is called,
	SchemeBuilder.Register(addKnownTypes)
}

func addKnownTypes(scheme *runtime.Scheme) error {
	// Add the types P4 and P4List to scheme
	scheme.AddKnownTypes(SchemeGroupVersion, &P4{}, &P4List{})
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
