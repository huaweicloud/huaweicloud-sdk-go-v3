/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// 更新floatingip的请求体
type NeutronUpdateFloatingIpRequestBody struct {
	Floatingip *UpdateFloatingIpOption `json:"floatingip"`
}
