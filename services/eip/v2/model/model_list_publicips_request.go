/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// Request Object
type ListPublicipsRequest struct {
	Marker string `json:"marker,omitempty"`
	Limit int32 `json:"limit,omitempty"`
	IpVersion int32 `json:"ip_version,omitempty"`
	EnterpriseProjectId string `json:"enterprise_project_id,omitempty"`
	PortId string `json:"port_id,omitempty"`
	PublicIpAddress string `json:"public_ip_address,omitempty"`
	PrivateIpAddress string `json:"private_ip_address,omitempty"`
	Id string `json:"id,omitempty"`
}
