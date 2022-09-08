package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExtendedConfigs struct {

	// 扩展参数属性名称
	Name *string `json:"name,omitempty"`

	// cdm-base64编码后的值
	Value *string `json:"value,omitempty"`
}

func (o ExtendedConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtendedConfigs struct{}"
	}

	return strings.Join([]string{"ExtendedConfigs", string(data)}, " ")
}
