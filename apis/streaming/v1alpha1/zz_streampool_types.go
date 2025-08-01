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

type CustomEncryptionKeyInitParameters struct {

	// (Updatable) Custom Encryption Key (Master Key) ocid.
	// +crossplane:generate:reference:type=github.com/oracle/provider-oci/apis/kms/v1alpha1.Key
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`
}

type CustomEncryptionKeyObservation struct {

	// (Updatable) Custom Encryption Key (Master Key) ocid.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Life cycle State of the custom key
	KeyState *string `json:"keyState,omitempty" tf:"key_state,omitempty"`
}

type CustomEncryptionKeyParameters struct {

	// (Updatable) Custom Encryption Key (Master Key) ocid.
	// +crossplane:generate:reference:type=github.com/oracle/provider-oci/apis/kms/v1alpha1.Key
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`
}

type KafkaSettingsInitParameters struct {

	// (Updatable) Enable auto creation of topic on the server.
	AutoCreateTopicsEnable *bool `json:"autoCreateTopicsEnable,omitempty" tf:"auto_create_topics_enable,omitempty"`

	// (Updatable) The number of hours to keep a log file before deleting it (in hours).
	LogRetentionHours *float64 `json:"logRetentionHours,omitempty" tf:"log_retention_hours,omitempty"`

	// (Updatable) The default number of log partitions per topic.
	NumPartitions *float64 `json:"numPartitions,omitempty" tf:"num_partitions,omitempty"`
}

type KafkaSettingsObservation struct {

	// (Updatable) Enable auto creation of topic on the server.
	AutoCreateTopicsEnable *bool `json:"autoCreateTopicsEnable,omitempty" tf:"auto_create_topics_enable,omitempty"`

	// (Updatable) Bootstrap servers.
	BootstrapServers *string `json:"bootstrapServers,omitempty" tf:"bootstrap_servers,omitempty"`

	// (Updatable) The number of hours to keep a log file before deleting it (in hours).
	LogRetentionHours *float64 `json:"logRetentionHours,omitempty" tf:"log_retention_hours,omitempty"`

	// (Updatable) The default number of log partitions per topic.
	NumPartitions *float64 `json:"numPartitions,omitempty" tf:"num_partitions,omitempty"`
}

type KafkaSettingsParameters struct {

	// (Updatable) Enable auto creation of topic on the server.
	// +kubebuilder:validation:Optional
	AutoCreateTopicsEnable *bool `json:"autoCreateTopicsEnable,omitempty" tf:"auto_create_topics_enable,omitempty"`

	// (Updatable) The number of hours to keep a log file before deleting it (in hours).
	// +kubebuilder:validation:Optional
	LogRetentionHours *float64 `json:"logRetentionHours,omitempty" tf:"log_retention_hours,omitempty"`

	// (Updatable) The default number of log partitions per topic.
	// +kubebuilder:validation:Optional
	NumPartitions *float64 `json:"numPartitions,omitempty" tf:"num_partitions,omitempty"`
}

type PrivateEndpointSettingsInitParameters struct {

	// The optional list of network security groups to be used with the private endpoint of the stream pool. That value cannot be changed.
	// +crossplane:generate:reference:type=github.com/oracle/provider-oci/apis/core/v1alpha1.NetworkSecurityGroup
	// +listType=set
	NsgIds []*string `json:"nsgIds,omitempty" tf:"nsg_ids,omitempty"`

	// References to NetworkSecurityGroup in core to populate nsgIds.
	// +kubebuilder:validation:Optional
	NsgIdsRefs []v1.Reference `json:"nsgIdsRefs,omitempty" tf:"-"`

	// Selector for a list of NetworkSecurityGroup in core to populate nsgIds.
	// +kubebuilder:validation:Optional
	NsgIdsSelector *v1.Selector `json:"nsgIdsSelector,omitempty" tf:"-"`

	// The optional private IP you want to be associated with your private stream pool. That parameter can only be specified when the subnetId parameter is set. It cannot be changed. The private IP needs to be part of the CIDR range of the specified subnetId or the creation will fail. If not specified a random IP inside the subnet will be chosen. After the stream pool is created, a custom FQDN, pointing to this private IP, is created. The FQDN is then used to access the service instead of the private IP.
	// +crossplane:generate:reference:type=github.com/oracle/provider-oci/apis/core/v1alpha1.PrivateIP
	PrivateEndpointIP *string `json:"privateEndpointIp,omitempty" tf:"private_endpoint_ip,omitempty"`

	// Reference to a PrivateIP in core to populate privateEndpointIp.
	// +kubebuilder:validation:Optional
	PrivateEndpointIPRef *v1.Reference `json:"privateEndpointIpRef,omitempty" tf:"-"`

	// Selector for a PrivateIP in core to populate privateEndpointIp.
	// +kubebuilder:validation:Optional
	PrivateEndpointIPSelector *v1.Selector `json:"privateEndpointIpSelector,omitempty" tf:"-"`

	// If specified, the stream pool will be private and only accessible from inside that subnet. Producing-to and consuming-from a stream inside a private stream pool can also only be done from inside the subnet. That value cannot be changed.
	// +crossplane:generate:reference:type=github.com/oracle/provider-oci/apis/core/v1alpha1.Subnet
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in core to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in core to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

type PrivateEndpointSettingsObservation struct {

	// The optional list of network security groups to be used with the private endpoint of the stream pool. That value cannot be changed.
	// +listType=set
	NsgIds []*string `json:"nsgIds,omitempty" tf:"nsg_ids,omitempty"`

	// The optional private IP you want to be associated with your private stream pool. That parameter can only be specified when the subnetId parameter is set. It cannot be changed. The private IP needs to be part of the CIDR range of the specified subnetId or the creation will fail. If not specified a random IP inside the subnet will be chosen. After the stream pool is created, a custom FQDN, pointing to this private IP, is created. The FQDN is then used to access the service instead of the private IP.
	PrivateEndpointIP *string `json:"privateEndpointIp,omitempty" tf:"private_endpoint_ip,omitempty"`

	// If specified, the stream pool will be private and only accessible from inside that subnet. Producing-to and consuming-from a stream inside a private stream pool can also only be done from inside the subnet. That value cannot be changed.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type PrivateEndpointSettingsParameters struct {

	// The optional list of network security groups to be used with the private endpoint of the stream pool. That value cannot be changed.
	// +crossplane:generate:reference:type=github.com/oracle/provider-oci/apis/core/v1alpha1.NetworkSecurityGroup
	// +kubebuilder:validation:Optional
	// +listType=set
	NsgIds []*string `json:"nsgIds,omitempty" tf:"nsg_ids,omitempty"`

	// References to NetworkSecurityGroup in core to populate nsgIds.
	// +kubebuilder:validation:Optional
	NsgIdsRefs []v1.Reference `json:"nsgIdsRefs,omitempty" tf:"-"`

	// Selector for a list of NetworkSecurityGroup in core to populate nsgIds.
	// +kubebuilder:validation:Optional
	NsgIdsSelector *v1.Selector `json:"nsgIdsSelector,omitempty" tf:"-"`

	// The optional private IP you want to be associated with your private stream pool. That parameter can only be specified when the subnetId parameter is set. It cannot be changed. The private IP needs to be part of the CIDR range of the specified subnetId or the creation will fail. If not specified a random IP inside the subnet will be chosen. After the stream pool is created, a custom FQDN, pointing to this private IP, is created. The FQDN is then used to access the service instead of the private IP.
	// +crossplane:generate:reference:type=github.com/oracle/provider-oci/apis/core/v1alpha1.PrivateIP
	// +kubebuilder:validation:Optional
	PrivateEndpointIP *string `json:"privateEndpointIp,omitempty" tf:"private_endpoint_ip,omitempty"`

	// Reference to a PrivateIP in core to populate privateEndpointIp.
	// +kubebuilder:validation:Optional
	PrivateEndpointIPRef *v1.Reference `json:"privateEndpointIpRef,omitempty" tf:"-"`

	// Selector for a PrivateIP in core to populate privateEndpointIp.
	// +kubebuilder:validation:Optional
	PrivateEndpointIPSelector *v1.Selector `json:"privateEndpointIpSelector,omitempty" tf:"-"`

	// If specified, the stream pool will be private and only accessible from inside that subnet. Producing-to and consuming-from a stream inside a private stream pool can also only be done from inside the subnet. That value cannot be changed.
	// +crossplane:generate:reference:type=github.com/oracle/provider-oci/apis/core/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in core to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in core to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

type StreamPoolInitParameters struct {

	// (Updatable) The OCID of the compartment that contains the stream.
	// +crossplane:generate:reference:type=github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// Reference to a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDRef *v1.Reference `json:"compartmentIdRef,omitempty" tf:"-"`

	// Selector for a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDSelector *v1.Selector `json:"compartmentIdSelector,omitempty" tf:"-"`

	// (Updatable) The OCID of the custom encryption key to be used or deleted if currently being used.
	CustomEncryptionKey []CustomEncryptionKeyInitParameters `json:"customEncryptionKey,omitempty" tf:"custom_encryption_key,omitempty"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.  Example: {"Operations.CostCenter": "42"}
	// +mapType=granular
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair that is applied with no predefined name, type, or namespace. Exists for cross-compatibility only. For more information, see Resource Tags.  Example: {"Department": "Finance"}
	// +mapType=granular
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// (Updatable) Settings for the Kafka compatibility layer.
	KafkaSettings []KafkaSettingsInitParameters `json:"kafkaSettings,omitempty" tf:"kafka_settings,omitempty"`

	// (Updatable) The name of the stream pool. Avoid entering confidential information.  Example: MyStreamPool
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Optional parameters if a private stream pool is requested.
	PrivateEndpointSettings []PrivateEndpointSettingsInitParameters `json:"privateEndpointSettings,omitempty" tf:"private_endpoint_settings,omitempty"`
}

type StreamPoolObservation struct {

	// (Updatable) The OCID of the compartment that contains the stream.
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// (Updatable) The OCID of the custom encryption key to be used or deleted if currently being used.
	CustomEncryptionKey []CustomEncryptionKeyObservation `json:"customEncryptionKey,omitempty" tf:"custom_encryption_key,omitempty"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.  Example: {"Operations.CostCenter": "42"}
	// +mapType=granular
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// The FQDN used to access the streams inside the stream pool (same FQDN as the messagesEndpoint attribute of a Stream object). If the stream pool is private, the FQDN is customized and can only be accessed from inside the associated subnetId, otherwise the FQDN is publicly resolvable. Depending on which protocol you attempt to use, you need to either prepend https or append the Kafka port.
	EndpointFqdn *string `json:"endpointFqdn,omitempty" tf:"endpoint_fqdn,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair that is applied with no predefined name, type, or namespace. Exists for cross-compatibility only. For more information, see Resource Tags.  Example: {"Department": "Finance"}
	// +mapType=granular
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// The OCID of the stream pool.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// True if the stream pool is private, false otherwise. The associated endpoint and subnetId of a private stream pool can be retrieved through the GetStreamPool API.
	IsPrivate *bool `json:"isPrivate,omitempty" tf:"is_private,omitempty"`

	// (Updatable) Settings for the Kafka compatibility layer.
	KafkaSettings []KafkaSettingsObservation `json:"kafkaSettings,omitempty" tf:"kafka_settings,omitempty"`

	// Any additional details about the current state of the stream.
	LifecycleStateDetails *string `json:"lifecycleStateDetails,omitempty" tf:"lifecycle_state_details,omitempty"`

	// (Updatable) The name of the stream pool. Avoid entering confidential information.  Example: MyStreamPool
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Optional parameters if a private stream pool is requested.
	PrivateEndpointSettings []PrivateEndpointSettingsObservation `json:"privateEndpointSettings,omitempty" tf:"private_endpoint_settings,omitempty"`

	// The current state of the stream pool.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The date and time the stream pool was created, expressed in in RFC 3339 timestamp format.  Example: 2018-04-20T00:00:07.405Z
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`
}

type StreamPoolParameters struct {

	// (Updatable) The OCID of the compartment that contains the stream.
	// +crossplane:generate:reference:type=github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment
	// +kubebuilder:validation:Optional
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// Reference to a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDRef *v1.Reference `json:"compartmentIdRef,omitempty" tf:"-"`

	// Selector for a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDSelector *v1.Selector `json:"compartmentIdSelector,omitempty" tf:"-"`

	// (Updatable) The OCID of the custom encryption key to be used or deleted if currently being used.
	// +kubebuilder:validation:Optional
	CustomEncryptionKey []CustomEncryptionKeyParameters `json:"customEncryptionKey,omitempty" tf:"custom_encryption_key,omitempty"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.  Example: {"Operations.CostCenter": "42"}
	// +kubebuilder:validation:Optional
	// +mapType=granular
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair that is applied with no predefined name, type, or namespace. Exists for cross-compatibility only. For more information, see Resource Tags.  Example: {"Department": "Finance"}
	// +kubebuilder:validation:Optional
	// +mapType=granular
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// (Updatable) Settings for the Kafka compatibility layer.
	// +kubebuilder:validation:Optional
	KafkaSettings []KafkaSettingsParameters `json:"kafkaSettings,omitempty" tf:"kafka_settings,omitempty"`

	// (Updatable) The name of the stream pool. Avoid entering confidential information.  Example: MyStreamPool
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Optional parameters if a private stream pool is requested.
	// +kubebuilder:validation:Optional
	PrivateEndpointSettings []PrivateEndpointSettingsParameters `json:"privateEndpointSettings,omitempty" tf:"private_endpoint_settings,omitempty"`
}

// StreamPoolSpec defines the desired state of StreamPool
type StreamPoolSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StreamPoolParameters `json:"forProvider"`
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
	InitProvider StreamPoolInitParameters `json:"initProvider,omitempty"`
}

// StreamPoolStatus defines the observed state of StreamPool.
type StreamPoolStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StreamPoolObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// StreamPool is the Schema for the StreamPools API. Provides the Stream Pool resource in Oracle Cloud Infrastructure Streaming service
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type StreamPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   StreamPoolSpec   `json:"spec"`
	Status StreamPoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StreamPoolList contains a list of StreamPools
type StreamPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StreamPool `json:"items"`
}

// Repository type metadata.
var (
	StreamPool_Kind             = "StreamPool"
	StreamPool_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StreamPool_Kind}.String()
	StreamPool_KindAPIVersion   = StreamPool_Kind + "." + CRDGroupVersion.String()
	StreamPool_GroupVersionKind = CRDGroupVersion.WithKind(StreamPool_Kind)
)

func init() {
	SchemeBuilder.Register(&StreamPool{}, &StreamPoolList{})
}
