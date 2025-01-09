package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubnetBandwidthsResponse Response Object
type ListSubnetBandwidthsResponse struct {

	// 云办公带宽信息。
	Bandwidths *[]SubnetBandwidthDetail `json:"bandwidths,omitempty"`

	// 总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSubnetBandwidthsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubnetBandwidthsResponse struct{}"
	}

	return strings.Join([]string{"ListSubnetBandwidthsResponse", string(data)}, " ")
}
