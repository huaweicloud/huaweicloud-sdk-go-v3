/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// 创建floatingip对象
type NeutronCreateFloatingIpRequestBody struct {
	Floatingip *CreateFloatingIpOption `json:"floatingip"`
}
