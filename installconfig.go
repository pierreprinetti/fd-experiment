package openstack

// MachinePool stores the configuration for a machine pool installed
// on OpenStack.
type MachinePool struct {
	// FlavorName defines the OpenStack Nova flavor.
	// eg. m1.large
	FlavorName string `json:"type"`

	// RootVolume defines the root volume for instances in the machine pool.
	// The instances use ephemeral disks if not set.
	// +optional
	RootVolume *RootVolume `json:"rootVolume,omitempty"`

	// AdditionalNetworkIDs contains IDs of additional networks for machines,
	// where each ID is presented in UUID v4 format.
	// Allowed address pairs won't be created for the additional networks.
	// +optional
	AdditionalNetworkIDs []string `json:"additionalNetworkIDs,omitempty"`

	// AdditionalSecurityGroupIDs contains IDs of additional security groups for machines,
	// where each ID is presented in UUID v4 format.
	// +optional
	AdditionalSecurityGroupIDs []string `json:"additionalSecurityGroupIDs,omitempty"`

	// ServerGroupPolicy will be used to create the Server Group that will contain all the machines of this MachinePool.
	// Defaults to "soft-anti-affinity".
	ServerGroupPolicy ServerGroupPolicy `json:"serverGroupPolicy,omitempty"`

	// Zones is the list of availability zones where the instances should be deployed.
	// If no zones are provided, all instances will be deployed on OpenStack Nova default availability zone
	// +optional
	Zones []string `json:"zones,omitempty"`

	// NEW FIELDS:
	Ports          []PortOpts `json:"ports,omitempty"`
	FailureDomains []FailureDomain
}

type FailureDomain struct {
	// ComputeAvailabilityZone is the name of a valid nova availability zone. The server will be created in this availability zone.
	// +optional
	ComputeAvailabilityZone string `json:"computeAvailabilityZone,omitempty"`

	// StorageAvailabilityZone is the name of a valid cinder availability
	// zone. This will be the availability zone of the root volume if one is
	// specified.
	// +optional
	StorageAvailabilityZone string `json:"storageAvailabilityZone,omitempty"`

	// Ports defines a set of port targets which can be referenced by machines in this failure domain.
	// +optional
	PortTargets []NamedPortTarget `json:"portTargets,omitempty"`
}

type PortOpts struct {
	// targetID references a portTarget in failureDomain. If targetID is
	// specified, both network and fixedIPs must be empty.
	TargetID string `json:"targetID,omitempty"`

	PortTarget `json:",inline"`

	DisablePortSecurity *bool    `json:"disablePortSecurity,omitempty"`
	SecurityGroups      []string `json:"securityGroups,omitempty"`
	Tags                []string `json:"tags,omitempty"`
}

type PortTarget struct {
	// Network is a query for an openstack network that the port will be created or discovered on.
	// This will fail if the query returns more than one network.
	Network *NetworkFilter `json:"network,omitempty"`
	// Specify pairs of subnet and/or IP address. These should be subnets of the network with the given NetworkID.
	FixedIPs []FixedIP `json:"fixedIPs,omitempty"`
}

type FixedIP struct {
	// subnetID specifies the ID of the subnet where the fixed IP will be allocated.
	SubnetID string `json:"subnetID"`
	// ipAddress is a specific IP address to use in the given subnet. Port
	// creation will fail if the address is not available. If not specified,
	// an available IP from the given subnet will be selected automatically.
	IPAddress string `json:"ipAddress,omitempty"`
}

type NamedPortTarget struct {
	// kubebuilder:validation:Required
	ID         string `json:"id"`
	PortTarget `json:",inline"`
}

type NetworkFilter struct {
	Name string `json:"name,omitempty"`
	ID   string `json:"id,omitempty"`
}
