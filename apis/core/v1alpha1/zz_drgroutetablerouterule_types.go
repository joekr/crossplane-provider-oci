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

type DrgRouteTableRouteRuleInitParameters struct {

	// (Updatable) This is the range of IP addresses used for matching when routing traffic. Only CIDR_BLOCK values are allowed.
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`

	// Type of destination for the rule. Required if direction = EGRESS. Allowed values:
	DestinationType *string `json:"destinationType,omitempty" tf:"destination_type,omitempty"`

	// The OCID of the DRG route table.
	// +crossplane:generate:reference:type=DrgRouteTable
	DrgRouteTableID *string `json:"drgRouteTableId,omitempty" tf:"drg_route_table_id,omitempty"`

	// Reference to a DrgRouteTable to populate drgRouteTableId.
	// +kubebuilder:validation:Optional
	DrgRouteTableIDRef *v1.Reference `json:"drgRouteTableIdRef,omitempty" tf:"-"`

	// Selector for a DrgRouteTable to populate drgRouteTableId.
	// +kubebuilder:validation:Optional
	DrgRouteTableIDSelector *v1.Selector `json:"drgRouteTableIdSelector,omitempty" tf:"-"`

	// The OCID of the next hop DRG attachment. The next hop DRG attachment is responsible for reaching the network destination.
	// +crossplane:generate:reference:type=DrgAttachment
	NextHopDrgAttachmentID *string `json:"nextHopDrgAttachmentId,omitempty" tf:"next_hop_drg_attachment_id,omitempty"`

	// Reference to a DrgAttachment to populate nextHopDrgAttachmentId.
	// +kubebuilder:validation:Optional
	NextHopDrgAttachmentIDRef *v1.Reference `json:"nextHopDrgAttachmentIdRef,omitempty" tf:"-"`

	// Selector for a DrgAttachment to populate nextHopDrgAttachmentId.
	// +kubebuilder:validation:Optional
	NextHopDrgAttachmentIDSelector *v1.Selector `json:"nextHopDrgAttachmentIdSelector,omitempty" tf:"-"`
}

type DrgRouteTableRouteRuleObservation struct {

	// Additional properties for the route, computed by the service.
	// +mapType=granular
	Attributes map[string]*string `json:"attributes,omitempty" tf:"attributes,omitempty"`

	// (Updatable) This is the range of IP addresses used for matching when routing traffic. Only CIDR_BLOCK values are allowed.
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`

	// Type of destination for the rule. Required if direction = EGRESS. Allowed values:
	DestinationType *string `json:"destinationType,omitempty" tf:"destination_type,omitempty"`

	// The OCID of the DRG route table.
	DrgRouteTableID *string `json:"drgRouteTableId,omitempty" tf:"drg_route_table_id,omitempty"`

	// The Oracle-assigned ID of the DRG route rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Indicates that if the next hop attachment does not exist, so traffic for this route is discarded without notification.
	IsBlackhole *bool `json:"isBlackhole,omitempty" tf:"is_blackhole,omitempty"`

	// Indicates that the route was not imported due to a conflict between route rules.
	IsConflict *bool `json:"isConflict,omitempty" tf:"is_conflict,omitempty"`

	// The OCID of the next hop DRG attachment. The next hop DRG attachment is responsible for reaching the network destination.
	NextHopDrgAttachmentID *string `json:"nextHopDrgAttachmentId,omitempty" tf:"next_hop_drg_attachment_id,omitempty"`

	// The earliest origin of a route. If a route is advertised to a DRG through an IPsec tunnel attachment, and is propagated to peered DRGs via RPC attachments, the route's provenance in the peered DRGs remains IPSEC_TUNNEL, because that is the earliest origin.
	RouteProvenance *string `json:"routeProvenance,omitempty" tf:"route_provenance,omitempty"`

	// You can specify static routes for the DRG route table using the API. The DRG learns dynamic routes from the DRG attachments using various routing protocols.
	RouteType *string `json:"routeType,omitempty" tf:"route_type,omitempty"`
}

type DrgRouteTableRouteRuleParameters struct {

	// (Updatable) This is the range of IP addresses used for matching when routing traffic. Only CIDR_BLOCK values are allowed.
	// +kubebuilder:validation:Optional
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`

	// Type of destination for the rule. Required if direction = EGRESS. Allowed values:
	// +kubebuilder:validation:Optional
	DestinationType *string `json:"destinationType,omitempty" tf:"destination_type,omitempty"`

	// The OCID of the DRG route table.
	// +crossplane:generate:reference:type=DrgRouteTable
	// +kubebuilder:validation:Optional
	DrgRouteTableID *string `json:"drgRouteTableId,omitempty" tf:"drg_route_table_id,omitempty"`

	// Reference to a DrgRouteTable to populate drgRouteTableId.
	// +kubebuilder:validation:Optional
	DrgRouteTableIDRef *v1.Reference `json:"drgRouteTableIdRef,omitempty" tf:"-"`

	// Selector for a DrgRouteTable to populate drgRouteTableId.
	// +kubebuilder:validation:Optional
	DrgRouteTableIDSelector *v1.Selector `json:"drgRouteTableIdSelector,omitempty" tf:"-"`

	// The OCID of the next hop DRG attachment. The next hop DRG attachment is responsible for reaching the network destination.
	// +crossplane:generate:reference:type=DrgAttachment
	// +kubebuilder:validation:Optional
	NextHopDrgAttachmentID *string `json:"nextHopDrgAttachmentId,omitempty" tf:"next_hop_drg_attachment_id,omitempty"`

	// Reference to a DrgAttachment to populate nextHopDrgAttachmentId.
	// +kubebuilder:validation:Optional
	NextHopDrgAttachmentIDRef *v1.Reference `json:"nextHopDrgAttachmentIdRef,omitempty" tf:"-"`

	// Selector for a DrgAttachment to populate nextHopDrgAttachmentId.
	// +kubebuilder:validation:Optional
	NextHopDrgAttachmentIDSelector *v1.Selector `json:"nextHopDrgAttachmentIdSelector,omitempty" tf:"-"`
}

// DrgRouteTableRouteRuleSpec defines the desired state of DrgRouteTableRouteRule
type DrgRouteTableRouteRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DrgRouteTableRouteRuleParameters `json:"forProvider"`
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
	InitProvider DrgRouteTableRouteRuleInitParameters `json:"initProvider,omitempty"`
}

// DrgRouteTableRouteRuleStatus defines the observed state of DrgRouteTableRouteRule.
type DrgRouteTableRouteRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DrgRouteTableRouteRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DrgRouteTableRouteRule is the Schema for the DrgRouteTableRouteRules API. Provides the Drg Route Table Route Rule resource in Oracle Cloud Infrastructure Core service
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type DrgRouteTableRouteRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.destination) || (has(self.initProvider) && has(self.initProvider.destination))",message="spec.forProvider.destination is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.destinationType) || (has(self.initProvider) && has(self.initProvider.destinationType))",message="spec.forProvider.destinationType is a required parameter"
	Spec   DrgRouteTableRouteRuleSpec   `json:"spec"`
	Status DrgRouteTableRouteRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DrgRouteTableRouteRuleList contains a list of DrgRouteTableRouteRules
type DrgRouteTableRouteRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DrgRouteTableRouteRule `json:"items"`
}

// Repository type metadata.
var (
	DrgRouteTableRouteRule_Kind             = "DrgRouteTableRouteRule"
	DrgRouteTableRouteRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DrgRouteTableRouteRule_Kind}.String()
	DrgRouteTableRouteRule_KindAPIVersion   = DrgRouteTableRouteRule_Kind + "." + CRDGroupVersion.String()
	DrgRouteTableRouteRule_GroupVersionKind = CRDGroupVersion.WithKind(DrgRouteTableRouteRule_Kind)
)

func init() {
	SchemeBuilder.Register(&DrgRouteTableRouteRule{}, &DrgRouteTableRouteRuleList{})
}
