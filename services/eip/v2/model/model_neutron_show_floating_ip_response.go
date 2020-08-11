/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// Response Object
type NeutronShowFloatingIpResponse struct {
	Floatingip *FloatingIpResp `json:"floatingip,omitempty"`
}
