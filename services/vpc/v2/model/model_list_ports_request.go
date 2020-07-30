/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// Request Object
type ListPortsRequest struct {
	Name string `json:"name,omitempty"`
	Id string `json:"id,omitempty"`
	Limit int32 `json:"limit,omitempty"`
	AdminStateUp bool `json:"admin_state_up,omitempty"`
	NetworkId string `json:"network_id,omitempty"`
	MacAddress string `json:"mac_address,omitempty"`
	DeviceId string `json:"device_id,omitempty"`
	DeviceOwner string `json:"device_owner,omitempty"`
	Status string `json:"status,omitempty"`
	Marker string `json:"marker,omitempty"`
	FixedIps string `json:"fixed_ips,omitempty"`
	EnterpriseProjectId string `json:"enterprise_project_id,omitempty"`
}
