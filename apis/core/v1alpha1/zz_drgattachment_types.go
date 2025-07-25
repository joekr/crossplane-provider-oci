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

type DrgAttachmentInitParameters struct {

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.  Example: {"Operations.CostCenter": "42"}
	// +mapType=granular
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The OCID of the DRG.
	// +crossplane:generate:reference:type=Drg
	DrgID *string `json:"drgId,omitempty" tf:"drg_id,omitempty"`

	// Reference to a Drg to populate drgId.
	// +kubebuilder:validation:Optional
	DrgIDRef *v1.Reference `json:"drgIdRef,omitempty" tf:"-"`

	// Selector for a Drg to populate drgId.
	// +kubebuilder:validation:Optional
	DrgIDSelector *v1.Selector `json:"drgIdSelector,omitempty" tf:"-"`

	// (Updatable) The OCID of the DRG route table that is assigned to this attachment.
	// +crossplane:generate:reference:type=DrgRouteTable
	DrgRouteTableID *string `json:"drgRouteTableId,omitempty" tf:"drg_route_table_id,omitempty"`

	// Reference to a DrgRouteTable to populate drgRouteTableId.
	// +kubebuilder:validation:Optional
	DrgRouteTableIDRef *v1.Reference `json:"drgRouteTableIdRef,omitempty" tf:"-"`

	// Selector for a DrgRouteTable to populate drgRouteTableId.
	// +kubebuilder:validation:Optional
	DrgRouteTableIDSelector *v1.Selector `json:"drgRouteTableIdSelector,omitempty" tf:"-"`

	// (Updatable) The OCID of the export route distribution used to specify how routes in the assigned DRG route table are advertised to the attachment. If this value is null, no routes are advertised through this attachment.
	// This field cannot be set by the user while creating the resource and gets a default value on creation. This can be only be updated to its default value. If this fields needs to be set to null, remove_export_drg_route_distribution_trigger needs to be used.
	ExportDrgRouteDistributionID *string `json:"exportDrgRouteDistributionId,omitempty" tf:"export_drg_route_distribution_id,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags.  Example: {"Department": "Finance"}
	// +mapType=granular
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// (Updatable)
	NetworkDetails []NetworkDetailsInitParameters `json:"networkDetails,omitempty" tf:"network_details,omitempty"`

	// (Updatable) An optional property when set to true during update disables the export of route Distribution by setting export_drg_route_distribution_id to null.
	RemoveExportDrgRouteDistributionTrigger *bool `json:"removeExportDrgRouteDistributionTrigger,omitempty" tf:"remove_export_drg_route_distribution_trigger,omitempty"`

	// (Updatable) This is the OCID of the route table that is used to route the traffic as it enters a VCN through this attachment.
	RouteTableID *string `json:"routeTableId,omitempty" tf:"route_table_id,omitempty"`

	// (Deprecated)  The OCID of the VCN. This field is deprecated. Instead, use the networkDetails field to specify the OCID of the attached resource.
	VcnID *string `json:"vcnId,omitempty" tf:"vcn_id,omitempty"`
}

type DrgAttachmentObservation struct {

	// The OCID of the compartment containing the DRG attachment.
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.  Example: {"Operations.CostCenter": "42"}
	// +mapType=granular
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The OCID of the DRG.
	DrgID *string `json:"drgId,omitempty" tf:"drg_id,omitempty"`

	// (Updatable) The OCID of the DRG route table that is assigned to this attachment.
	DrgRouteTableID *string `json:"drgRouteTableId,omitempty" tf:"drg_route_table_id,omitempty"`

	// (Updatable) The OCID of the export route distribution used to specify how routes in the assigned DRG route table are advertised to the attachment. If this value is null, no routes are advertised through this attachment.
	// This field cannot be set by the user while creating the resource and gets a default value on creation. This can be only be updated to its default value. If this fields needs to be set to null, remove_export_drg_route_distribution_trigger needs to be used.
	ExportDrgRouteDistributionID *string `json:"exportDrgRouteDistributionId,omitempty" tf:"export_drg_route_distribution_id,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags.  Example: {"Department": "Finance"}
	// +mapType=granular
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// The OCID of the network attached to the DRG.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Indicates whether the DRG attachment and attached network live in a different tenancy than the DRG.  Example: false
	IsCrossTenancy *bool `json:"isCrossTenancy,omitempty" tf:"is_cross_tenancy,omitempty"`

	// (Updatable)
	NetworkDetails []NetworkDetailsObservation `json:"networkDetails,omitempty" tf:"network_details,omitempty"`

	// (Updatable) An optional property when set to true during update disables the export of route Distribution by setting export_drg_route_distribution_id to null.
	RemoveExportDrgRouteDistributionTrigger *bool `json:"removeExportDrgRouteDistributionTrigger,omitempty" tf:"remove_export_drg_route_distribution_trigger,omitempty"`

	// (Updatable) This is the OCID of the route table that is used to route the traffic as it enters a VCN through this attachment.
	RouteTableID *string `json:"routeTableId,omitempty" tf:"route_table_id,omitempty"`

	// The DRG attachment's current state.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The date and time the DRG attachment was created, in the format defined by RFC3339.  Example: 2016-08-25T21:10:29.600Z
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	// (Deprecated)  The OCID of the VCN. This field is deprecated. Instead, use the networkDetails field to specify the OCID of the attached resource.
	VcnID *string `json:"vcnId,omitempty" tf:"vcn_id,omitempty"`
}

type DrgAttachmentParameters struct {

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.  Example: {"Operations.CostCenter": "42"}
	// +kubebuilder:validation:Optional
	// +mapType=granular
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The OCID of the DRG.
	// +crossplane:generate:reference:type=Drg
	// +kubebuilder:validation:Optional
	DrgID *string `json:"drgId,omitempty" tf:"drg_id,omitempty"`

	// Reference to a Drg to populate drgId.
	// +kubebuilder:validation:Optional
	DrgIDRef *v1.Reference `json:"drgIdRef,omitempty" tf:"-"`

	// Selector for a Drg to populate drgId.
	// +kubebuilder:validation:Optional
	DrgIDSelector *v1.Selector `json:"drgIdSelector,omitempty" tf:"-"`

	// (Updatable) The OCID of the DRG route table that is assigned to this attachment.
	// +crossplane:generate:reference:type=DrgRouteTable
	// +kubebuilder:validation:Optional
	DrgRouteTableID *string `json:"drgRouteTableId,omitempty" tf:"drg_route_table_id,omitempty"`

	// Reference to a DrgRouteTable to populate drgRouteTableId.
	// +kubebuilder:validation:Optional
	DrgRouteTableIDRef *v1.Reference `json:"drgRouteTableIdRef,omitempty" tf:"-"`

	// Selector for a DrgRouteTable to populate drgRouteTableId.
	// +kubebuilder:validation:Optional
	DrgRouteTableIDSelector *v1.Selector `json:"drgRouteTableIdSelector,omitempty" tf:"-"`

	// (Updatable) The OCID of the export route distribution used to specify how routes in the assigned DRG route table are advertised to the attachment. If this value is null, no routes are advertised through this attachment.
	// This field cannot be set by the user while creating the resource and gets a default value on creation. This can be only be updated to its default value. If this fields needs to be set to null, remove_export_drg_route_distribution_trigger needs to be used.
	// +kubebuilder:validation:Optional
	ExportDrgRouteDistributionID *string `json:"exportDrgRouteDistributionId,omitempty" tf:"export_drg_route_distribution_id,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags.  Example: {"Department": "Finance"}
	// +kubebuilder:validation:Optional
	// +mapType=granular
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// (Updatable)
	// +kubebuilder:validation:Optional
	NetworkDetails []NetworkDetailsParameters `json:"networkDetails,omitempty" tf:"network_details,omitempty"`

	// (Updatable) An optional property when set to true during update disables the export of route Distribution by setting export_drg_route_distribution_id to null.
	// +kubebuilder:validation:Optional
	RemoveExportDrgRouteDistributionTrigger *bool `json:"removeExportDrgRouteDistributionTrigger,omitempty" tf:"remove_export_drg_route_distribution_trigger,omitempty"`

	// (Updatable) This is the OCID of the route table that is used to route the traffic as it enters a VCN through this attachment.
	// +kubebuilder:validation:Optional
	RouteTableID *string `json:"routeTableId,omitempty" tf:"route_table_id,omitempty"`

	// (Deprecated)  The OCID of the VCN. This field is deprecated. Instead, use the networkDetails field to specify the OCID of the attached resource.
	// +kubebuilder:validation:Optional
	VcnID *string `json:"vcnId,omitempty" tf:"vcn_id,omitempty"`
}

type NetworkDetailsInitParameters struct {

	// The OCID of the network attached to the DRG.
	// +crossplane:generate:reference:type=Vcn
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Reference to a Vcn to populate id.
	// +kubebuilder:validation:Optional
	IDRef *v1.Reference `json:"idRef,omitempty" tf:"-"`

	// Selector for a Vcn to populate id.
	// +kubebuilder:validation:Optional
	IDSelector *v1.Selector `json:"idSelector,omitempty" tf:"-"`

	// (Updatable) This is the OCID of the route table that is used to route the traffic as it enters a VCN through this attachment.
	// +crossplane:generate:reference:type=RouteTable
	RouteTableID *string `json:"routeTableId,omitempty" tf:"route_table_id,omitempty"`

	// Reference to a RouteTable to populate routeTableId.
	// +kubebuilder:validation:Optional
	RouteTableIDRef *v1.Reference `json:"routeTableIdRef,omitempty" tf:"-"`

	// Selector for a RouteTable to populate routeTableId.
	// +kubebuilder:validation:Optional
	RouteTableIDSelector *v1.Selector `json:"routeTableIdSelector,omitempty" tf:"-"`

	// (Updatable) The type can be one of these values: IPSEC_TUNNEL, REMOTE_PEERING_CONNECTION, VCN, VIRTUAL_CIRCUIT
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (Updatable) Indicates whether the VCN CIDRs or the individual subnet CIDRs are imported from the attachment. Routes from the VCN ingress route table are always imported.
	VcnRouteType *string `json:"vcnRouteType,omitempty" tf:"vcn_route_type,omitempty"`
}

type NetworkDetailsObservation struct {

	// The OCID of the network attached to the DRG.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Ids []*string `json:"ids,omitempty" tf:"ids,omitempty"`

	// The IPSec connection that contains the attached IPSec tunnel.
	IpsecConnectionID *string `json:"ipsecConnectionId,omitempty" tf:"ipsec_connection_id,omitempty"`

	// (Updatable) This is the OCID of the route table that is used to route the traffic as it enters a VCN through this attachment.
	RouteTableID *string `json:"routeTableId,omitempty" tf:"route_table_id,omitempty"`

	// The OCID of the network attached to the DRG.
	TransportAttachmentID *string `json:"transportAttachmentId,omitempty" tf:"transport_attachment_id,omitempty"`

	TransportOnlyMode *bool `json:"transportOnlyMode,omitempty" tf:"transport_only_mode,omitempty"`

	// (Updatable) The type can be one of these values: IPSEC_TUNNEL, REMOTE_PEERING_CONNECTION, VCN, VIRTUAL_CIRCUIT
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (Updatable) Indicates whether the VCN CIDRs or the individual subnet CIDRs are imported from the attachment. Routes from the VCN ingress route table are always imported.
	VcnRouteType *string `json:"vcnRouteType,omitempty" tf:"vcn_route_type,omitempty"`
}

type NetworkDetailsParameters struct {

	// The OCID of the network attached to the DRG.
	// +crossplane:generate:reference:type=Vcn
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Reference to a Vcn to populate id.
	// +kubebuilder:validation:Optional
	IDRef *v1.Reference `json:"idRef,omitempty" tf:"-"`

	// Selector for a Vcn to populate id.
	// +kubebuilder:validation:Optional
	IDSelector *v1.Selector `json:"idSelector,omitempty" tf:"-"`

	// (Updatable) This is the OCID of the route table that is used to route the traffic as it enters a VCN through this attachment.
	// +crossplane:generate:reference:type=RouteTable
	// +kubebuilder:validation:Optional
	RouteTableID *string `json:"routeTableId,omitempty" tf:"route_table_id,omitempty"`

	// Reference to a RouteTable to populate routeTableId.
	// +kubebuilder:validation:Optional
	RouteTableIDRef *v1.Reference `json:"routeTableIdRef,omitempty" tf:"-"`

	// Selector for a RouteTable to populate routeTableId.
	// +kubebuilder:validation:Optional
	RouteTableIDSelector *v1.Selector `json:"routeTableIdSelector,omitempty" tf:"-"`

	// (Updatable) The type can be one of these values: IPSEC_TUNNEL, REMOTE_PEERING_CONNECTION, VCN, VIRTUAL_CIRCUIT
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`

	// (Updatable) Indicates whether the VCN CIDRs or the individual subnet CIDRs are imported from the attachment. Routes from the VCN ingress route table are always imported.
	// +kubebuilder:validation:Optional
	VcnRouteType *string `json:"vcnRouteType,omitempty" tf:"vcn_route_type,omitempty"`
}

// DrgAttachmentSpec defines the desired state of DrgAttachment
type DrgAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DrgAttachmentParameters `json:"forProvider"`
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
	InitProvider DrgAttachmentInitParameters `json:"initProvider,omitempty"`
}

// DrgAttachmentStatus defines the observed state of DrgAttachment.
type DrgAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DrgAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DrgAttachment is the Schema for the DrgAttachments API. Provides the Drg Attachment resource in Oracle Cloud Infrastructure Core service
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type DrgAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DrgAttachmentSpec   `json:"spec"`
	Status            DrgAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DrgAttachmentList contains a list of DrgAttachments
type DrgAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DrgAttachment `json:"items"`
}

// Repository type metadata.
var (
	DrgAttachment_Kind             = "DrgAttachment"
	DrgAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DrgAttachment_Kind}.String()
	DrgAttachment_KindAPIVersion   = DrgAttachment_Kind + "." + CRDGroupVersion.String()
	DrgAttachment_GroupVersionKind = CRDGroupVersion.WithKind(DrgAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&DrgAttachment{}, &DrgAttachmentList{})
}
