package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateGroupBandwidthInfoRequest struct {

	// 分片ID。
	GroupId string `json:"group_id"`

	// 目标带宽（Mbit/s）。
	TargetBandwidth int32 `json:"target_bandwidth"`
}

func (o UpdateGroupBandwidthInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGroupBandwidthInfoRequest struct{}"
	}

	return strings.Join([]string{"UpdateGroupBandwidthInfoRequest", string(data)}, " ")
}
