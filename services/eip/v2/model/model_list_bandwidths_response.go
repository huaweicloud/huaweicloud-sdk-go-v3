/*
 * EIP
 *
 * 云服务接口
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListBandwidthsResponse struct {
	// 带宽列表对象
	Bandwidths *[]BandwidthResp `json:"bandwidths,omitempty"`
}

func (o ListBandwidthsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListBandwidthsResponse", string(data)}, " ")
}
