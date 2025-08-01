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

type SubnetInitParameters struct {

	// Controls whether the subnet is regional or specific to an availability domain. Oracle recommends creating regional subnets because they're more flexible and make it easier to implement failover across availability domains. Originally, AD-specific subnets were the only kind available to use.
	AvailabilityDomain *string `json:"availabilityDomain,omitempty" tf:"availability_domain,omitempty"`

	// (Updatable) The CIDR IP address range of the subnet. The CIDR must maintain the following rules -
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// (Updatable) The OCID of the compartment to contain the subnet.
	// +crossplane:generate:reference:type=github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// Reference to a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDRef *v1.Reference `json:"compartmentIdRef,omitempty" tf:"-"`

	// Selector for a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDSelector *v1.Selector `json:"compartmentIdSelector,omitempty" tf:"-"`

	// (Updatable) The OCID of the set of DHCP options the subnet will use. If you don't provide a value, the subnet uses the VCN's default set of DHCP options.
	// +crossplane:generate:reference:type=DHCPOptions
	DHCPOptionsID *string `json:"dhcpOptionsId,omitempty" tf:"dhcp_options_id,omitempty"`

	// Reference to a DHCPOptions to populate dhcpOptionsId.
	// +kubebuilder:validation:Optional
	DHCPOptionsIDRef *v1.Reference `json:"dhcpOptionsIdRef,omitempty" tf:"-"`

	// Selector for a DHCPOptions to populate dhcpOptionsId.
	// +kubebuilder:validation:Optional
	DHCPOptionsIDSelector *v1.Selector `json:"dhcpOptionsIdSelector,omitempty" tf:"-"`

	// A DNS label for the subnet, used in conjunction with the VNIC's hostname and VCN's DNS label to form a fully qualified domain name (FQDN) for each VNIC within this subnet (for example, bminstance1.subnet123.vcn1.oraclevcn.com). Must be an alphanumeric string that begins with a letter and is unique within the VCN. The value cannot be changed.
	DNSLabel *string `json:"dnsLabel,omitempty" tf:"dns_label,omitempty"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.  Example: {"Operations.CostCenter": "42"}
	// +mapType=granular
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags.  Example: {"Department": "Finance"}
	// +mapType=granular
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// (Updatable) Use this to enable IPv6 addressing for this subnet. The VCN must be enabled for IPv6. You can't change this subnet characteristic later. All subnets are /64 in size. The subnet portion of the IPv6 address is the fourth hextet from the left (1111 in the following example).
	Ipv6CidrBlock *string `json:"ipv6cidrBlock,omitempty" tf:"ipv6cidr_block,omitempty"`

	// (Updatable) The list of all IPv6 CIDR blocks (Oracle allocated IPv6 GUA, ULA or private IPv6 CIDR blocks, BYOIPv6 CIDR blocks) for the subnet that meets the following criteria:
	Ipv6CidrBlocks []*string `json:"ipv6cidrBlocks,omitempty" tf:"ipv6cidr_blocks,omitempty"`

	// Whether to disallow ingress internet traffic to VNICs within this subnet. Defaults to false.
	ProhibitInternetIngress *bool `json:"prohibitInternetIngress,omitempty" tf:"prohibit_internet_ingress,omitempty"`

	// Whether VNICs within this subnet can have public IP addresses. Defaults to false, which means VNICs created in this subnet will automatically be assigned public IP addresses unless specified otherwise during instance launch or VNIC creation (with the assignPublicIp flag in CreateVnicDetails). If prohibitPublicIpOnVnic is set to true, VNICs created in this subnet cannot have public IP addresses (that is, it's a private subnet).
	ProhibitPublicIPOnVnic *bool `json:"prohibitPublicIpOnVnic,omitempty" tf:"prohibit_public_ip_on_vnic,omitempty"`

	// (Updatable) The OCID of the route table the subnet will use. If you don't provide a value, the subnet uses the VCN's default route table.
	// +crossplane:generate:reference:type=RouteTable
	RouteTableID *string `json:"routeTableId,omitempty" tf:"route_table_id,omitempty"`

	// Reference to a RouteTable to populate routeTableId.
	// +kubebuilder:validation:Optional
	RouteTableIDRef *v1.Reference `json:"routeTableIdRef,omitempty" tf:"-"`

	// Selector for a RouteTable to populate routeTableId.
	// +kubebuilder:validation:Optional
	RouteTableIDSelector *v1.Selector `json:"routeTableIdSelector,omitempty" tf:"-"`

	// References to SecurityList to populate securityListIds.
	// +kubebuilder:validation:Optional
	SecurityListIDRefs []v1.Reference `json:"securityListIdRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityList to populate securityListIds.
	// +kubebuilder:validation:Optional
	SecurityListIDSelector *v1.Selector `json:"securityListIdSelector,omitempty" tf:"-"`

	// (Updatable) The OCIDs of the security list or lists the subnet will use. If you don't provide a value, the subnet uses the VCN's default security list. Remember that security lists are associated with the subnet, but the rules are applied to the individual VNICs in the subnet.
	// +crossplane:generate:reference:type=SecurityList
	// +crossplane:generate:reference:refFieldName=SecurityListIDRefs
	// +crossplane:generate:reference:selectorFieldName=SecurityListIDSelector
	// +listType=set
	SecurityListIds []*string `json:"securityListIds,omitempty" tf:"security_list_ids,omitempty"`

	// The OCID of the VCN to contain the subnet.
	// +crossplane:generate:reference:type=Vcn
	VcnID *string `json:"vcnId,omitempty" tf:"vcn_id,omitempty"`

	// Reference to a Vcn to populate vcnId.
	// +kubebuilder:validation:Optional
	VcnIDRef *v1.Reference `json:"vcnIdRef,omitempty" tf:"-"`

	// Selector for a Vcn to populate vcnId.
	// +kubebuilder:validation:Optional
	VcnIDSelector *v1.Selector `json:"vcnIdSelector,omitempty" tf:"-"`
}

type SubnetObservation struct {

	// Controls whether the subnet is regional or specific to an availability domain. Oracle recommends creating regional subnets because they're more flexible and make it easier to implement failover across availability domains. Originally, AD-specific subnets were the only kind available to use.
	AvailabilityDomain *string `json:"availabilityDomain,omitempty" tf:"availability_domain,omitempty"`

	// (Updatable) The CIDR IP address range of the subnet. The CIDR must maintain the following rules -
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// (Updatable) The OCID of the compartment to contain the subnet.
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// (Updatable) The OCID of the set of DHCP options the subnet will use. If you don't provide a value, the subnet uses the VCN's default set of DHCP options.
	DHCPOptionsID *string `json:"dhcpOptionsId,omitempty" tf:"dhcp_options_id,omitempty"`

	// A DNS label for the subnet, used in conjunction with the VNIC's hostname and VCN's DNS label to form a fully qualified domain name (FQDN) for each VNIC within this subnet (for example, bminstance1.subnet123.vcn1.oraclevcn.com). Must be an alphanumeric string that begins with a letter and is unique within the VCN. The value cannot be changed.
	DNSLabel *string `json:"dnsLabel,omitempty" tf:"dns_label,omitempty"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.  Example: {"Operations.CostCenter": "42"}
	// +mapType=granular
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags.  Example: {"Department": "Finance"}
	// +mapType=granular
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// The subnet's Oracle ID (OCID).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Updatable) Use this to enable IPv6 addressing for this subnet. The VCN must be enabled for IPv6. You can't change this subnet characteristic later. All subnets are /64 in size. The subnet portion of the IPv6 address is the fourth hextet from the left (1111 in the following example).
	Ipv6CidrBlock *string `json:"ipv6cidrBlock,omitempty" tf:"ipv6cidr_block,omitempty"`

	// (Updatable) The list of all IPv6 CIDR blocks (Oracle allocated IPv6 GUA, ULA or private IPv6 CIDR blocks, BYOIPv6 CIDR blocks) for the subnet that meets the following criteria:
	Ipv6CidrBlocks []*string `json:"ipv6cidrBlocks,omitempty" tf:"ipv6cidr_blocks,omitempty"`

	// For an IPv6-enabled subnet, this is the IPv6 address of the virtual router.  Example: 2001:0db8:0123:1111:89ab:cdef:1234:5678
	Ipv6VirtualRouterIP *string `json:"ipv6virtualRouterIp,omitempty" tf:"ipv6virtual_router_ip,omitempty"`

	// Whether to disallow ingress internet traffic to VNICs within this subnet. Defaults to false.
	ProhibitInternetIngress *bool `json:"prohibitInternetIngress,omitempty" tf:"prohibit_internet_ingress,omitempty"`

	// Whether VNICs within this subnet can have public IP addresses. Defaults to false, which means VNICs created in this subnet will automatically be assigned public IP addresses unless specified otherwise during instance launch or VNIC creation (with the assignPublicIp flag in CreateVnicDetails). If prohibitPublicIpOnVnic is set to true, VNICs created in this subnet cannot have public IP addresses (that is, it's a private subnet).
	ProhibitPublicIPOnVnic *bool `json:"prohibitPublicIpOnVnic,omitempty" tf:"prohibit_public_ip_on_vnic,omitempty"`

	// (Updatable) The OCID of the route table the subnet will use. If you don't provide a value, the subnet uses the VCN's default route table.
	RouteTableID *string `json:"routeTableId,omitempty" tf:"route_table_id,omitempty"`

	// (Updatable) The OCIDs of the security list or lists the subnet will use. If you don't provide a value, the subnet uses the VCN's default security list. Remember that security lists are associated with the subnet, but the rules are applied to the individual VNICs in the subnet.
	// +listType=set
	SecurityListIds []*string `json:"securityListIds,omitempty" tf:"security_list_ids,omitempty"`

	// The subnet's current state.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The subnet's domain name, which consists of the subnet's DNS label, the VCN's DNS label, and the oraclevcn.com domain.
	SubnetDomainName *string `json:"subnetDomainName,omitempty" tf:"subnet_domain_name,omitempty"`

	// The date and time the subnet was created, in the format defined by RFC3339.  Example: 2016-08-25T21:10:29.600Z
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	// The OCID of the VCN to contain the subnet.
	VcnID *string `json:"vcnId,omitempty" tf:"vcn_id,omitempty"`

	// The IP address of the virtual router.  Example: 10.0.14.1
	VirtualRouterIP *string `json:"virtualRouterIp,omitempty" tf:"virtual_router_ip,omitempty"`

	// The MAC address of the virtual router.  Example: 00:00:00:00:00:01
	VirtualRouterMac *string `json:"virtualRouterMac,omitempty" tf:"virtual_router_mac,omitempty"`
}

type SubnetParameters struct {

	// Controls whether the subnet is regional or specific to an availability domain. Oracle recommends creating regional subnets because they're more flexible and make it easier to implement failover across availability domains. Originally, AD-specific subnets were the only kind available to use.
	// +kubebuilder:validation:Optional
	AvailabilityDomain *string `json:"availabilityDomain,omitempty" tf:"availability_domain,omitempty"`

	// (Updatable) The CIDR IP address range of the subnet. The CIDR must maintain the following rules -
	// +kubebuilder:validation:Optional
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// (Updatable) The OCID of the compartment to contain the subnet.
	// +crossplane:generate:reference:type=github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment
	// +kubebuilder:validation:Optional
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// Reference to a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDRef *v1.Reference `json:"compartmentIdRef,omitempty" tf:"-"`

	// Selector for a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDSelector *v1.Selector `json:"compartmentIdSelector,omitempty" tf:"-"`

	// (Updatable) The OCID of the set of DHCP options the subnet will use. If you don't provide a value, the subnet uses the VCN's default set of DHCP options.
	// +crossplane:generate:reference:type=DHCPOptions
	// +kubebuilder:validation:Optional
	DHCPOptionsID *string `json:"dhcpOptionsId,omitempty" tf:"dhcp_options_id,omitempty"`

	// Reference to a DHCPOptions to populate dhcpOptionsId.
	// +kubebuilder:validation:Optional
	DHCPOptionsIDRef *v1.Reference `json:"dhcpOptionsIdRef,omitempty" tf:"-"`

	// Selector for a DHCPOptions to populate dhcpOptionsId.
	// +kubebuilder:validation:Optional
	DHCPOptionsIDSelector *v1.Selector `json:"dhcpOptionsIdSelector,omitempty" tf:"-"`

	// A DNS label for the subnet, used in conjunction with the VNIC's hostname and VCN's DNS label to form a fully qualified domain name (FQDN) for each VNIC within this subnet (for example, bminstance1.subnet123.vcn1.oraclevcn.com). Must be an alphanumeric string that begins with a letter and is unique within the VCN. The value cannot be changed.
	// +kubebuilder:validation:Optional
	DNSLabel *string `json:"dnsLabel,omitempty" tf:"dns_label,omitempty"`

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

	// (Updatable) Use this to enable IPv6 addressing for this subnet. The VCN must be enabled for IPv6. You can't change this subnet characteristic later. All subnets are /64 in size. The subnet portion of the IPv6 address is the fourth hextet from the left (1111 in the following example).
	// +kubebuilder:validation:Optional
	Ipv6CidrBlock *string `json:"ipv6cidrBlock,omitempty" tf:"ipv6cidr_block,omitempty"`

	// (Updatable) The list of all IPv6 CIDR blocks (Oracle allocated IPv6 GUA, ULA or private IPv6 CIDR blocks, BYOIPv6 CIDR blocks) for the subnet that meets the following criteria:
	// +kubebuilder:validation:Optional
	Ipv6CidrBlocks []*string `json:"ipv6cidrBlocks,omitempty" tf:"ipv6cidr_blocks,omitempty"`

	// Whether to disallow ingress internet traffic to VNICs within this subnet. Defaults to false.
	// +kubebuilder:validation:Optional
	ProhibitInternetIngress *bool `json:"prohibitInternetIngress,omitempty" tf:"prohibit_internet_ingress,omitempty"`

	// Whether VNICs within this subnet can have public IP addresses. Defaults to false, which means VNICs created in this subnet will automatically be assigned public IP addresses unless specified otherwise during instance launch or VNIC creation (with the assignPublicIp flag in CreateVnicDetails). If prohibitPublicIpOnVnic is set to true, VNICs created in this subnet cannot have public IP addresses (that is, it's a private subnet).
	// +kubebuilder:validation:Optional
	ProhibitPublicIPOnVnic *bool `json:"prohibitPublicIpOnVnic,omitempty" tf:"prohibit_public_ip_on_vnic,omitempty"`

	// (Updatable) The OCID of the route table the subnet will use. If you don't provide a value, the subnet uses the VCN's default route table.
	// +crossplane:generate:reference:type=RouteTable
	// +kubebuilder:validation:Optional
	RouteTableID *string `json:"routeTableId,omitempty" tf:"route_table_id,omitempty"`

	// Reference to a RouteTable to populate routeTableId.
	// +kubebuilder:validation:Optional
	RouteTableIDRef *v1.Reference `json:"routeTableIdRef,omitempty" tf:"-"`

	// Selector for a RouteTable to populate routeTableId.
	// +kubebuilder:validation:Optional
	RouteTableIDSelector *v1.Selector `json:"routeTableIdSelector,omitempty" tf:"-"`

	// References to SecurityList to populate securityListIds.
	// +kubebuilder:validation:Optional
	SecurityListIDRefs []v1.Reference `json:"securityListIdRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityList to populate securityListIds.
	// +kubebuilder:validation:Optional
	SecurityListIDSelector *v1.Selector `json:"securityListIdSelector,omitempty" tf:"-"`

	// (Updatable) The OCIDs of the security list or lists the subnet will use. If you don't provide a value, the subnet uses the VCN's default security list. Remember that security lists are associated with the subnet, but the rules are applied to the individual VNICs in the subnet.
	// +crossplane:generate:reference:type=SecurityList
	// +crossplane:generate:reference:refFieldName=SecurityListIDRefs
	// +crossplane:generate:reference:selectorFieldName=SecurityListIDSelector
	// +kubebuilder:validation:Optional
	// +listType=set
	SecurityListIds []*string `json:"securityListIds,omitempty" tf:"security_list_ids,omitempty"`

	// The OCID of the VCN to contain the subnet.
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

// SubnetSpec defines the desired state of Subnet
type SubnetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubnetParameters `json:"forProvider"`
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
	InitProvider SubnetInitParameters `json:"initProvider,omitempty"`
}

// SubnetStatus defines the observed state of Subnet.
type SubnetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubnetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Subnet is the Schema for the Subnets API. Provides the Subnet resource in Oracle Cloud Infrastructure Core service
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type Subnet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.cidrBlock) || (has(self.initProvider) && has(self.initProvider.cidrBlock))",message="spec.forProvider.cidrBlock is a required parameter"
	Spec   SubnetSpec   `json:"spec"`
	Status SubnetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubnetList contains a list of Subnets
type SubnetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Subnet `json:"items"`
}

// Repository type metadata.
var (
	Subnet_Kind             = "Subnet"
	Subnet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Subnet_Kind}.String()
	Subnet_KindAPIVersion   = Subnet_Kind + "." + CRDGroupVersion.String()
	Subnet_GroupVersionKind = CRDGroupVersion.WithKind(Subnet_Kind)
)

func init() {
	SchemeBuilder.Register(&Subnet{}, &SubnetList{})
}
