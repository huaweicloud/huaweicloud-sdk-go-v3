/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// Request Object
type NeutronUpdateFloatingIpRequest struct {
	FloatingipId string `json:"floatingip_id"`
	Body *NeutronUpdateFloatingIpRequestBody `json:"body,omitempty"`
}
