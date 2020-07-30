/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// Response Object
type BatchCreateSubNetworkInterfaceV3Response struct {
	// 请求ID
	RequestId string `json:"request_id,omitempty"`
	// 批量创建辅助弹性网卡的响应体
	SubNetworkInterfaces []SubNetworkInterface `json:"sub_network_interfaces,omitempty"`
}
