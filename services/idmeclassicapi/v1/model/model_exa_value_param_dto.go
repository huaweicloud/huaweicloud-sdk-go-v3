package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExaValueParamDto struct {

	// 中文名称。
	Name *string `json:"name,omitempty"`

	// 值。
	Value *interface{} `json:"value,omitempty"`
}

func (o ExaValueParamDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExaValueParamDto struct{}"
	}

	return strings.Join([]string{"ExaValueParamDto", string(data)}, " ")
}
