package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateInstanceTopicReqTopicOtherConfigs struct {

	// 配置名称
	Name *string `json:"name,omitempty"`

	// 配置值
	Value *string `json:"value,omitempty"`
}

func (o UpdateInstanceTopicReqTopicOtherConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceTopicReqTopicOtherConfigs struct{}"
	}

	return strings.Join([]string{"UpdateInstanceTopicReqTopicOtherConfigs", string(data)}, " ")
}
