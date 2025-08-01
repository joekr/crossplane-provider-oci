/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RouteTableAttachmentInitParameters struct {

	// The OCID of the route table.
	RouteTableID *string `json:"routeTableId,omitempty" tf:"route_table_id,omitempty"`

	// The OCID of the subnet.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type RouteTableAttachmentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The OCID of the route table.
	RouteTableID *string `json:"routeTableId,omitempty" tf:"route_table_id,omitempty"`

	// The OCID of the subnet.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type RouteTableAttachmentParameters struct {

	// The OCID of the route table.
	// +kubebuilder:validation:Optional
	RouteTableID *string `json:"routeTableId,omitempty" tf:"route_table_id,omitempty"`

	// The OCID of the subnet.
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

// RouteTableAttachmentSpec defines the desired state of RouteTableAttachment
type RouteTableAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RouteTableAttachmentParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider RouteTableAttachmentInitParameters `json:"initProvider,omitempty"`
}

// RouteTableAttachmentStatus defines the observed state of RouteTableAttachment.
type RouteTableAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RouteTableAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RouteTableAttachment is the Schema for the RouteTableAttachments API. Provides the ability to associate a route table for a subnet in Oracle Cloud Infrastructure Core service
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type RouteTableAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.routeTableId) || (has(self.initProvider) && has(self.initProvider.routeTableId))",message="spec.forProvider.routeTableId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.subnetId) || (has(self.initProvider) && has(self.initProvider.subnetId))",message="spec.forProvider.subnetId is a required parameter"
	Spec   RouteTableAttachmentSpec   `json:"spec"`
	Status RouteTableAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouteTableAttachmentList contains a list of RouteTableAttachments
type RouteTableAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouteTableAttachment `json:"items"`
}

// Repository type metadata.
var (
	RouteTableAttachment_Kind             = "RouteTableAttachment"
	RouteTableAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RouteTableAttachment_Kind}.String()
	RouteTableAttachment_KindAPIVersion   = RouteTableAttachment_Kind + "." + CRDGroupVersion.String()
	RouteTableAttachment_GroupVersionKind = CRDGroupVersion.WithKind(RouteTableAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&RouteTableAttachment{}, &RouteTableAttachmentList{})
}
