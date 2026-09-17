package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMissingIndexSwitchRequest Request Object
type ShowMissingIndexSwitchRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 引擎类型
	EngineType *string `json:"engine_type,omitempty"`
}

func (o ShowMissingIndexSwitchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMissingIndexSwitchRequest struct{}"
	}

	return strings.Join([]string{"ShowMissingIndexSwitchRequest", string(data)}, " ")
}
