package network

// Address represents an IP address
type Address struct {
	Addr      string
	PrefixLen int
}

// IPAM represents IP Address Management
type IPAM struct {
	Driver string
	Config []IPAMConfig
}

// IPAMConfig represents IPAM configurations
type IPAMConfig struct {
	Subnet     string            `json:",omitempty"`
	IPRange    string            `json:",omitempty"`
	Gateway    string            `json:",omitempty"`
	AuxAddress map[string]string `json:"AuxiliaryAddresses,omitempty"`
}

// EndpointIPAMConfig represents IPAM configurations for the endpoint
type EndpointIPAMConfig struct {
	IPv4Address string `json:",omitempty"`
	IPv6Address string `json:",omitempty"`
}

// EndpointSettings stores the network endpoint details
type EndpointSettings struct {
	// Configurations
	IPAMConfig *EndpointIPAMConfig
	// Operational data
	EndpointID          string
	Gateway             string
	IPAddress           string
	IPPrefixLen         int
	IPv6Gateway         string
	GlobalIPv6Address   string
	GlobalIPv6PrefixLen int
	MacAddress          string
	Aliases             []string
}

// CleanOperationalData resets the operational data for this endpoint
func (es *EndpointSettings) CleanOperationalData() {
	es.EndpointID = ""
	es.Gateway = ""
	es.IPAddress = ""
	es.IPPrefixLen = 0
	es.IPv6Gateway = ""
	es.GlobalIPv6Address = ""
	es.GlobalIPv6PrefixLen = 0
	es.MacAddress = ""
}

// NetworkingConfig represents the container's networking configuration for each of its interfaces
// Carries the networink configs specified in the `docker run` and `docker network connect` commands
type NetworkingConfig struct {
	EndpointsConfig map[string]*EndpointSettings // Endpoint configs for each conencting network
}

// HasUserDefinedIPAddress returns whether the passed endpoint configuration contains IP address configuration
func HasUserDefinedIPAddress(epConfig *EndpointSettings) bool {
	return epConfig != nil && epConfig.IPAMConfig != nil && (len(epConfig.IPAMConfig.IPv4Address) > 0 || len(epConfig.IPAMConfig.IPv6Address) > 0)
}
