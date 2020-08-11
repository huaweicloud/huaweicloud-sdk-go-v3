/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// Request Object
type NeutronDeleteFloatingIpRequest struct {
	FloatingipId string `json:"floatingip_id"`
}
