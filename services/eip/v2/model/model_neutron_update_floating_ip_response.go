/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// Response Object
type NeutronUpdateFloatingIpResponse struct {
	Floatingip *PostAndPutFloatingIpResp `json:"floatingip,omitempty"`
}
