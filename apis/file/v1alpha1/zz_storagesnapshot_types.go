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

type StorageSnapshotInitParameters struct {

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: {"Operations.CostCenter": "42"}
	// +mapType=granular
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	ExpirationTime *string `json:"expirationTime,omitempty" tf:"expiration_time,omitempty"`

	// The OCID of the file system to take a snapshot of.
	// +crossplane:generate:reference:type=StorageFileSystem
	FileSystemID *string `json:"fileSystemId,omitempty" tf:"file_system_id,omitempty"`

	// Reference to a StorageFileSystem to populate fileSystemId.
	// +kubebuilder:validation:Optional
	FileSystemIDRef *v1.Reference `json:"fileSystemIdRef,omitempty" tf:"-"`

	// Selector for a StorageFileSystem to populate fileSystemId.
	// +kubebuilder:validation:Optional
	FileSystemIDSelector *v1.Selector `json:"fileSystemIdSelector,omitempty" tf:"-"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags. Example: {"Department": "Finance"}
	// +mapType=granular
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	IsLockOverride *bool `json:"isLockOverride,omitempty" tf:"is_lock_override,omitempty"`

	Locks []StorageSnapshotLocksInitParameters `json:"locks,omitempty" tf:"locks,omitempty"`

	// Name of the snapshot. This value is immutable. It must also be unique with respect to all other non-DELETED snapshots on the associated file system.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type StorageSnapshotLocksInitParameters struct {
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// The OCID of the snapshot.
	RelatedResourceID *string `json:"relatedResourceId,omitempty" tf:"related_resource_id,omitempty"`

	// The date and time the snapshot was created, expressed in RFC 3339 timestamp format.  Example: 2016-08-25T21:10:29.600Z
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type StorageSnapshotLocksObservation struct {
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// The OCID of the snapshot.
	RelatedResourceID *string `json:"relatedResourceId,omitempty" tf:"related_resource_id,omitempty"`

	// The date and time the snapshot was created, expressed in RFC 3339 timestamp format.  Example: 2016-08-25T21:10:29.600Z
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type StorageSnapshotLocksParameters struct {

	// +kubebuilder:validation:Optional
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// The OCID of the snapshot.
	// +kubebuilder:validation:Optional
	RelatedResourceID *string `json:"relatedResourceId,omitempty" tf:"related_resource_id,omitempty"`

	// The date and time the snapshot was created, expressed in RFC 3339 timestamp format.  Example: 2016-08-25T21:10:29.600Z
	// +kubebuilder:validation:Optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type StorageSnapshotObservation struct {

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: {"Operations.CostCenter": "42"}
	// +mapType=granular
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	ExpirationTime *string `json:"expirationTime,omitempty" tf:"expiration_time,omitempty"`

	// The OCID of the file system to take a snapshot of.
	FileSystemID *string `json:"fileSystemId,omitempty" tf:"file_system_id,omitempty"`

	// The OCID of the snapshot.
	FilesystemSnapshotPolicyID *string `json:"filesystemSnapshotPolicyId,omitempty" tf:"filesystem_snapshot_policy_id,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags. Example: {"Department": "Finance"}
	// +mapType=granular
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// The OCID of the snapshot.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies whether the snapshot has been cloned. See Cloning a File System.
	IsCloneSource *bool `json:"isCloneSource,omitempty" tf:"is_clone_source,omitempty"`

	IsLockOverride *bool `json:"isLockOverride,omitempty" tf:"is_lock_override,omitempty"`

	// Additional information about the current lifecycleState.
	LifecycleDetails *string `json:"lifecycleDetails,omitempty" tf:"lifecycle_details,omitempty"`

	Locks []StorageSnapshotLocksObservation `json:"locks,omitempty" tf:"locks,omitempty"`

	// Name of the snapshot. This value is immutable. It must also be unique with respect to all other non-DELETED snapshots on the associated file system.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// An OCID identifying the parent from which this snapshot was cloned. If this snapshot was not cloned, then the provenanceId is the same as the snapshot id value. If this snapshot was cloned, then the provenanceId value is the parent's provenanceId. See Cloning a File System.
	ProvenanceID *string `json:"provenanceId,omitempty" tf:"provenance_id,omitempty"`

	// The date and time the snapshot was taken, expressed in RFC 3339 timestamp format. This value might be the same or different from timeCreated depending on the following factors:
	SnapshotTime *string `json:"snapshotTime,omitempty" tf:"snapshot_time,omitempty"`

	// Specifies the generation type of the snapshot.
	SnapshotType *string `json:"snapshotType,omitempty" tf:"snapshot_type,omitempty"`

	// The current state of the snapshot.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// +mapType=granular
	SystemTags map[string]*string `json:"systemTags,omitempty" tf:"system_tags,omitempty"`

	// The date and time the snapshot was created, expressed in RFC 3339 timestamp format.  Example: 2016-08-25T21:10:29.600Z
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`
}

type StorageSnapshotParameters struct {

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: {"Operations.CostCenter": "42"}
	// +kubebuilder:validation:Optional
	// +mapType=granular
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// +kubebuilder:validation:Optional
	ExpirationTime *string `json:"expirationTime,omitempty" tf:"expiration_time,omitempty"`

	// The OCID of the file system to take a snapshot of.
	// +crossplane:generate:reference:type=StorageFileSystem
	// +kubebuilder:validation:Optional
	FileSystemID *string `json:"fileSystemId,omitempty" tf:"file_system_id,omitempty"`

	// Reference to a StorageFileSystem to populate fileSystemId.
	// +kubebuilder:validation:Optional
	FileSystemIDRef *v1.Reference `json:"fileSystemIdRef,omitempty" tf:"-"`

	// Selector for a StorageFileSystem to populate fileSystemId.
	// +kubebuilder:validation:Optional
	FileSystemIDSelector *v1.Selector `json:"fileSystemIdSelector,omitempty" tf:"-"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags. Example: {"Department": "Finance"}
	// +kubebuilder:validation:Optional
	// +mapType=granular
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// +kubebuilder:validation:Optional
	IsLockOverride *bool `json:"isLockOverride,omitempty" tf:"is_lock_override,omitempty"`

	// +kubebuilder:validation:Optional
	Locks []StorageSnapshotLocksParameters `json:"locks,omitempty" tf:"locks,omitempty"`

	// Name of the snapshot. This value is immutable. It must also be unique with respect to all other non-DELETED snapshots on the associated file system.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// StorageSnapshotSpec defines the desired state of StorageSnapshot
type StorageSnapshotSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StorageSnapshotParameters `json:"forProvider"`
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
	InitProvider StorageSnapshotInitParameters `json:"initProvider,omitempty"`
}

// StorageSnapshotStatus defines the observed state of StorageSnapshot.
type StorageSnapshotStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StorageSnapshotObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// StorageSnapshot is the Schema for the StorageSnapshots API. Provides the Snapshot resource in Oracle Cloud Infrastructure File Storage service
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type StorageSnapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   StorageSnapshotSpec   `json:"spec"`
	Status StorageSnapshotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StorageSnapshotList contains a list of StorageSnapshots
type StorageSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageSnapshot `json:"items"`
}

// Repository type metadata.
var (
	StorageSnapshot_Kind             = "StorageSnapshot"
	StorageSnapshot_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StorageSnapshot_Kind}.String()
	StorageSnapshot_KindAPIVersion   = StorageSnapshot_Kind + "." + CRDGroupVersion.String()
	StorageSnapshot_GroupVersionKind = CRDGroupVersion.WithKind(StorageSnapshot_Kind)
)

func init() {
	SchemeBuilder.Register(&StorageSnapshot{}, &StorageSnapshotList{})
}
