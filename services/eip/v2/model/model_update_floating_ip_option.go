/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// 更新floatingip对象
type UpdateFloatingIpOption struct {
	// 端口id。
	PortId string `json:"port_id,omitempty"`
}
