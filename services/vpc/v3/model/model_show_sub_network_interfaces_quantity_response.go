/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// Response Object
type ShowSubNetworkInterfacesQuantityResponse struct {
	// 请求ID
	RequestId string `json:"request_id,omitempty"`
	// 辅助弹性网卡数目
	SubNetworkInterfaces int32 `json:"sub_network_interfaces,omitempty"`
}
