package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateBandwidthInfoRequest struct {

	// 分片带宽列表。
	GroupBandwidths *[]UpdateGroupBandwidthInfoRequest `json:"group_bandwidths,omitempty"`
}

func (o UpdateBandwidthInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBandwidthInfoRequest struct{}"
	}

	return strings.Join([]string{"UpdateBandwidthInfoRequest", string(data)}, " ")
}
