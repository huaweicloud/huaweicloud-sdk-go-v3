/*
 * BMS
 *
 * BMS Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// addresses数据结构说明
type AddressesInfo struct {
	// 裸金属服务器所属网络信息。key：表示裸金属服务器使用的虚拟私有云的ID。value：网络详细信息，具体请参见表4 address数据结构说明。
	VpcId *[]AddressInfo `json:"vpc_id,omitempty"`
}

func (o AddressesInfo) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"AddressesInfo", string(data)}, " ")
}
