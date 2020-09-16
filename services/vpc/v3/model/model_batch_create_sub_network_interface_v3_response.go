/*
 * VPC
 *
 * VPC Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchCreateSubNetworkInterfaceV3Response struct {
	// 请求ID
	RequestId *string `json:"request_id,omitempty"`
	// 批量创建辅助弹性网卡的响应体
	SubNetworkInterfaces *[]SubNetworkInterface `json:"sub_network_interfaces,omitempty"`
}

func (o BatchCreateSubNetworkInterfaceV3Response) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchCreateSubNetworkInterfaceV3Response", string(data)}, " ")
}
