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

type TagInitParameters struct {

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.  Example: {"Operations.CostCenter": "42"}
	// +mapType=granular
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) The description you assign to the tag during creation.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags.  Example: {"Department": "Finance"}
	// +mapType=granular
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// (Updatable) Indicates whether the tag is enabled for cost tracking.
	IsCostTracking *bool `json:"isCostTracking,omitempty" tf:"is_cost_tracking,omitempty"`

	// (Updatable) Indicates whether the tag is retired. See Retiring Key Definitions and Namespace Definitions.
	IsRetired *bool `json:"isRetired,omitempty" tf:"is_retired,omitempty"`

	// The name you assign to the tag during creation. This is the tag key definition. The name must be unique within the tag namespace and cannot be changed.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The OCID of the tag namespace.
	// +crossplane:generate:reference:type=TagNamespace
	TagNamespaceID *string `json:"tagNamespaceId,omitempty" tf:"tag_namespace_id,omitempty"`

	// Reference to a TagNamespace to populate tagNamespaceId.
	// +kubebuilder:validation:Optional
	TagNamespaceIDRef *v1.Reference `json:"tagNamespaceIdRef,omitempty" tf:"-"`

	// Selector for a TagNamespace to populate tagNamespaceId.
	// +kubebuilder:validation:Optional
	TagNamespaceIDSelector *v1.Selector `json:"tagNamespaceIdSelector,omitempty" tf:"-"`

	// is set, the user applying the tag to a resource can type in a static
	// value or leave the tag value empty.
	Validator []ValidatorInitParameters `json:"validator,omitempty" tf:"validator,omitempty"`
}

type TagObservation struct {

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.  Example: {"Operations.CostCenter": "42"}
	// +mapType=granular
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) The description you assign to the tag during creation.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags.  Example: {"Department": "Finance"}
	// +mapType=granular
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// The OCID of the tag definition.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Updatable) Indicates whether the tag is enabled for cost tracking.
	IsCostTracking *bool `json:"isCostTracking,omitempty" tf:"is_cost_tracking,omitempty"`

	// (Updatable) Indicates whether the tag is retired. See Retiring Key Definitions and Namespace Definitions.
	IsRetired *bool `json:"isRetired,omitempty" tf:"is_retired,omitempty"`

	// The name you assign to the tag during creation. This is the tag key definition. The name must be unique within the tag namespace and cannot be changed.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The tag's current state. After creating a tag, make sure its lifecycleState is ACTIVE before using it. After retiring a tag, make sure its lifecycleState is INACTIVE before using it. If you delete a tag, you cannot delete another tag until the deleted tag's lifecycleState changes from DELETING to DELETED.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The OCID of the tag namespace.
	TagNamespaceID *string `json:"tagNamespaceId,omitempty" tf:"tag_namespace_id,omitempty"`

	// Date and time the tag was created, in the format defined by RFC3339.  Example: 2016-08-25T21:10:29.600Z
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	// is set, the user applying the tag to a resource can type in a static
	// value or leave the tag value empty.
	Validator []ValidatorObservation `json:"validator,omitempty" tf:"validator,omitempty"`
}

type TagParameters struct {

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.  Example: {"Operations.CostCenter": "42"}
	// +kubebuilder:validation:Optional
	// +mapType=granular
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) The description you assign to the tag during creation.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags.  Example: {"Department": "Finance"}
	// +kubebuilder:validation:Optional
	// +mapType=granular
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// (Updatable) Indicates whether the tag is enabled for cost tracking.
	// +kubebuilder:validation:Optional
	IsCostTracking *bool `json:"isCostTracking,omitempty" tf:"is_cost_tracking,omitempty"`

	// (Updatable) Indicates whether the tag is retired. See Retiring Key Definitions and Namespace Definitions.
	// +kubebuilder:validation:Optional
	IsRetired *bool `json:"isRetired,omitempty" tf:"is_retired,omitempty"`

	// The name you assign to the tag during creation. This is the tag key definition. The name must be unique within the tag namespace and cannot be changed.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The OCID of the tag namespace.
	// +crossplane:generate:reference:type=TagNamespace
	// +kubebuilder:validation:Optional
	TagNamespaceID *string `json:"tagNamespaceId,omitempty" tf:"tag_namespace_id,omitempty"`

	// Reference to a TagNamespace to populate tagNamespaceId.
	// +kubebuilder:validation:Optional
	TagNamespaceIDRef *v1.Reference `json:"tagNamespaceIdRef,omitempty" tf:"-"`

	// Selector for a TagNamespace to populate tagNamespaceId.
	// +kubebuilder:validation:Optional
	TagNamespaceIDSelector *v1.Selector `json:"tagNamespaceIdSelector,omitempty" tf:"-"`

	// is set, the user applying the tag to a resource can type in a static
	// value or leave the tag value empty.
	// +kubebuilder:validation:Optional
	Validator []ValidatorParameters `json:"validator,omitempty" tf:"validator,omitempty"`
}

type ValidatorInitParameters struct {

	// (Updatable) Specifies the type of validation: a static value (no validation) or a list.
	ValidatorType *string `json:"validatorType,omitempty" tf:"validator_type,omitempty"`

	// (Applicable when validator_type=ENUM) (Updatable) The list of allowed values for a definedTag value.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type ValidatorObservation struct {

	// (Updatable) Specifies the type of validation: a static value (no validation) or a list.
	ValidatorType *string `json:"validatorType,omitempty" tf:"validator_type,omitempty"`

	// (Applicable when validator_type=ENUM) (Updatable) The list of allowed values for a definedTag value.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type ValidatorParameters struct {

	// (Updatable) Specifies the type of validation: a static value (no validation) or a list.
	// +kubebuilder:validation:Optional
	ValidatorType *string `json:"validatorType" tf:"validator_type,omitempty"`

	// (Applicable when validator_type=ENUM) (Updatable) The list of allowed values for a definedTag value.
	// +kubebuilder:validation:Optional
	Values []*string `json:"values" tf:"values,omitempty"`
}

// TagSpec defines the desired state of Tag
type TagSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TagParameters `json:"forProvider"`
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
	InitProvider TagInitParameters `json:"initProvider,omitempty"`
}

// TagStatus defines the observed state of Tag.
type TagStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TagObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Tag is the Schema for the Tags API. Provides the Tag resource in Oracle Cloud Infrastructure Identity service
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type Tag struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.description) || (has(self.initProvider) && has(self.initProvider.description))",message="spec.forProvider.description is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   TagSpec   `json:"spec"`
	Status TagStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TagList contains a list of Tags
type TagList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Tag `json:"items"`
}

// Repository type metadata.
var (
	Tag_Kind             = "Tag"
	Tag_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Tag_Kind}.String()
	Tag_KindAPIVersion   = Tag_Kind + "." + CRDGroupVersion.String()
	Tag_GroupVersionKind = CRDGroupVersion.WithKind(Tag_Kind)
)

func init() {
	SchemeBuilder.Register(&Tag{}, &TagList{})
}
