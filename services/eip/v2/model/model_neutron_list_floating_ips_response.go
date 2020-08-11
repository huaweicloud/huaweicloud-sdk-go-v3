/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// Response Object
type NeutronListFloatingIpsResponse struct {
	// floatingip对象列表
	Floatingips []FloatingIpResp `json:"floatingips,omitempty"`
	// marker分页结构
	FloatingipsLinks []Pager `json:"floatingips_links,omitempty"`
}
