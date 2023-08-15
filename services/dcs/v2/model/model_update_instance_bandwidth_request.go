package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceBandwidthRequest Request Object
type UpdateInstanceBandwidthRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o UpdateInstanceBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceBandwidthRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceBandwidthRequest", string(data)}, " ")
}
