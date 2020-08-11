/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// Request Object
type NeutronListFloatingIpsRequest struct {
	Limit string `json:"limit,omitempty"`
	Marker string `json:"marker,omitempty"`
	PageReverse bool `json:"page_reverse,omitempty"`
	Id string `json:"id,omitempty"`
	FloatingIpAddress string `json:"floating_ip_address,omitempty"`
	RouterId string `json:"router_id,omitempty"`
	PortId string `json:"port_id,omitempty"`
	FixedIpAddress string `json:"fixed_ip_address,omitempty"`
	TenantId string `json:"tenant_id,omitempty"`
	FloatingNetworkId string `json:"floating_network_id,omitempty"`
}
