/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// Response Object
type NeutronCreateFloatingIpResponse struct {
	Floatingip *PostAndPutFloatingIpResp `json:"floatingip,omitempty"`
}
