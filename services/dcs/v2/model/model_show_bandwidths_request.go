package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBandwidthsRequest Request Object
type ShowBandwidthsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ShowBandwidthsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBandwidthsRequest struct{}"
	}

	return strings.Join([]string{"ShowBandwidthsRequest", string(data)}, " ")
}
