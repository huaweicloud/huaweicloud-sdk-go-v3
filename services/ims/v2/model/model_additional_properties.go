package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AdditionalProperties 属性值
type AdditionalProperties struct {

	// 类型
	Type string `json:"type"`
}

func (o AdditionalProperties) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AdditionalProperties struct{}"
	}

	return strings.Join([]string{"AdditionalProperties", string(data)}, " ")
}
