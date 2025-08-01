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

type RouteRulesInitParameters struct {

	// (Updatable) Deprecated. Instead use destination and destinationType. Requests that include both cidrBlock and destination will be rejected.
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// (Updatable) An optional description of your choice for the rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Updatable) Conceptually, this is the range of IP addresses used for matching when routing traffic. Required if you provide a destinationType.
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`

	// (Updatable) Type of destination for the rule. Required if you provide a destination.
	DestinationType *string `json:"destinationType,omitempty" tf:"destination_type,omitempty"`

	// (Updatable) The OCID for the route rule's target. For information about the type of targets you can specify, see Route Tables.
	// +crossplane:generate:reference:type=NATGateway
	NetworkEntityID *string `json:"networkEntityId,omitempty" tf:"network_entity_id,omitempty"`

	// Reference to a NATGateway to populate networkEntityId.
	// +kubebuilder:validation:Optional
	NetworkEntityIDRef *v1.Reference `json:"networkEntityIdRef,omitempty" tf:"-"`

	// Selector for a NATGateway to populate networkEntityId.
	// +kubebuilder:validation:Optional
	NetworkEntityIDSelector *v1.Selector `json:"networkEntityIdSelector,omitempty" tf:"-"`

	// (Updatable) A route rule can be STATIC if manually added to the route table, LOCAL if added by Oracle Cloud Infrastructure to the route table.
	RouteType *string `json:"routeType,omitempty" tf:"route_type,omitempty"`
}

type RouteRulesObservation struct {

	// (Updatable) Deprecated. Instead use destination and destinationType. Requests that include both cidrBlock and destination will be rejected.
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// (Updatable) An optional description of your choice for the rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Updatable) Conceptually, this is the range of IP addresses used for matching when routing traffic. Required if you provide a destinationType.
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`

	// (Updatable) Type of destination for the rule. Required if you provide a destination.
	DestinationType *string `json:"destinationType,omitempty" tf:"destination_type,omitempty"`

	// (Updatable) The OCID for the route rule's target. For information about the type of targets you can specify, see Route Tables.
	NetworkEntityID *string `json:"networkEntityId,omitempty" tf:"network_entity_id,omitempty"`

	// (Updatable) A route rule can be STATIC if manually added to the route table, LOCAL if added by Oracle Cloud Infrastructure to the route table.
	RouteType *string `json:"routeType,omitempty" tf:"route_type,omitempty"`
}

type RouteRulesParameters struct {

	// (Updatable) Deprecated. Instead use destination and destinationType. Requests that include both cidrBlock and destination will be rejected.
	// +kubebuilder:validation:Optional
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// (Updatable) An optional description of your choice for the rule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Updatable) Conceptually, this is the range of IP addresses used for matching when routing traffic. Required if you provide a destinationType.
	// +kubebuilder:validation:Optional
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`

	// (Updatable) Type of destination for the rule. Required if you provide a destination.
	// +kubebuilder:validation:Optional
	DestinationType *string `json:"destinationType,omitempty" tf:"destination_type,omitempty"`

	// (Updatable) The OCID for the route rule's target. For information about the type of targets you can specify, see Route Tables.
	// +crossplane:generate:reference:type=NATGateway
	// +kubebuilder:validation:Optional
	NetworkEntityID *string `json:"networkEntityId,omitempty" tf:"network_entity_id,omitempty"`

	// Reference to a NATGateway to populate networkEntityId.
	// +kubebuilder:validation:Optional
	NetworkEntityIDRef *v1.Reference `json:"networkEntityIdRef,omitempty" tf:"-"`

	// Selector for a NATGateway to populate networkEntityId.
	// +kubebuilder:validation:Optional
	NetworkEntityIDSelector *v1.Selector `json:"networkEntityIdSelector,omitempty" tf:"-"`

	// (Updatable) A route rule can be STATIC if manually added to the route table, LOCAL if added by Oracle Cloud Infrastructure to the route table.
	// +kubebuilder:validation:Optional
	RouteType *string `json:"routeType,omitempty" tf:"route_type,omitempty"`
}

type RouteTableInitParameters struct {

	// (Updatable) The OCID of the compartment to contain the route table.
	// +crossplane:generate:reference:type=github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// Reference to a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDRef *v1.Reference `json:"compartmentIdRef,omitempty" tf:"-"`

	// Selector for a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDSelector *v1.Selector `json:"compartmentIdSelector,omitempty" tf:"-"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.  Example: {"Operations.CostCenter": "42"}
	// +mapType=granular
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags.  Example: {"Department": "Finance"}
	// +mapType=granular
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// (Updatable) The collection of rules used for routing destination IPs to network devices.
	RouteRules []RouteRulesInitParameters `json:"routeRules,omitempty" tf:"route_rules,omitempty"`

	// The OCID of the VCN the route table belongs to.
	// +crossplane:generate:reference:type=Vcn
	VcnID *string `json:"vcnId,omitempty" tf:"vcn_id,omitempty"`

	// Reference to a Vcn to populate vcnId.
	// +kubebuilder:validation:Optional
	VcnIDRef *v1.Reference `json:"vcnIdRef,omitempty" tf:"-"`

	// Selector for a Vcn to populate vcnId.
	// +kubebuilder:validation:Optional
	VcnIDSelector *v1.Selector `json:"vcnIdSelector,omitempty" tf:"-"`
}

type RouteTableObservation struct {

	// (Updatable) The OCID of the compartment to contain the route table.
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.  Example: {"Operations.CostCenter": "42"}
	// +mapType=granular
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags.  Example: {"Department": "Finance"}
	// +mapType=granular
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// The route table's Oracle ID (OCID).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Updatable) The collection of rules used for routing destination IPs to network devices.
	RouteRules []RouteRulesObservation `json:"routeRules,omitempty" tf:"route_rules,omitempty"`

	// The route table's current state.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The date and time the route table was created, in the format defined by RFC3339.  Example: 2016-08-25T21:10:29.600Z
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	// The OCID of the VCN the route table belongs to.
	VcnID *string `json:"vcnId,omitempty" tf:"vcn_id,omitempty"`
}

type RouteTableParameters struct {

	// (Updatable) The OCID of the compartment to contain the route table.
	// +crossplane:generate:reference:type=github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment
	// +kubebuilder:validation:Optional
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// Reference to a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDRef *v1.Reference `json:"compartmentIdRef,omitempty" tf:"-"`

	// Selector for a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDSelector *v1.Selector `json:"compartmentIdSelector,omitempty" tf:"-"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.  Example: {"Operations.CostCenter": "42"}
	// +kubebuilder:validation:Optional
	// +mapType=granular
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags.  Example: {"Department": "Finance"}
	// +kubebuilder:validation:Optional
	// +mapType=granular
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// (Updatable) The collection of rules used for routing destination IPs to network devices.
	// +kubebuilder:validation:Optional
	RouteRules []RouteRulesParameters `json:"routeRules,omitempty" tf:"route_rules,omitempty"`

	// The OCID of the VCN the route table belongs to.
	// +crossplane:generate:reference:type=Vcn
	// +kubebuilder:validation:Optional
	VcnID *string `json:"vcnId,omitempty" tf:"vcn_id,omitempty"`

	// Reference to a Vcn to populate vcnId.
	// +kubebuilder:validation:Optional
	VcnIDRef *v1.Reference `json:"vcnIdRef,omitempty" tf:"-"`

	// Selector for a Vcn to populate vcnId.
	// +kubebuilder:validation:Optional
	VcnIDSelector *v1.Selector `json:"vcnIdSelector,omitempty" tf:"-"`
}

// RouteTableSpec defines the desired state of RouteTable
type RouteTableSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RouteTableParameters `json:"forProvider"`
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
	InitProvider RouteTableInitParameters `json:"initProvider,omitempty"`
}

// RouteTableStatus defines the observed state of RouteTable.
type RouteTableStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RouteTableObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RouteTable is the Schema for the RouteTables API. Provides the Route Table resource in Oracle Cloud Infrastructure Core service
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type RouteTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouteTableSpec   `json:"spec"`
	Status            RouteTableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouteTableList contains a list of RouteTables
type RouteTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouteTable `json:"items"`
}

// Repository type metadata.
var (
	RouteTable_Kind             = "RouteTable"
	RouteTable_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RouteTable_Kind}.String()
	RouteTable_KindAPIVersion   = RouteTable_Kind + "." + CRDGroupVersion.String()
	RouteTable_GroupVersionKind = CRDGroupVersion.WithKind(RouteTable_Kind)
)

func init() {
	SchemeBuilder.Register(&RouteTable{}, &RouteTableList{})
}
