package v1alpha1

import (
	traefikv1alpha1 "github.com/fluxcd/flagger/pkg/apis/traefik/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TraefikService is the specification for a service (that an IngressRoute refers
// to) that is usually not a terminal service (i.e. not a pod of servers), as
// opposed to a Kubernetes Service. That is to say, it usually refers to other
// (children) services, which themselves can be TraefikServices or Services.
type TraefikService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec traefikv1alpha1.ServiceSpec `json:"spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TraefikServiceList is a list of TraefikService resources.
type TraefikServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []TraefikService `json:"items"`
}
