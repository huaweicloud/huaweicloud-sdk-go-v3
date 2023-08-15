package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceBiactiveRegionsRequest Request Object
type ShowInstanceBiactiveRegionsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ShowInstanceBiactiveRegionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceBiactiveRegionsRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceBiactiveRegionsRequest", string(data)}, " ")
}
