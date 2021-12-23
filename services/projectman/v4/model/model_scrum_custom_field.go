package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ScrumCustomField struct {
	// 自定义字段

	Name *string `json:"name,omitempty"`
	// 自定义字段值

	Value *string `json:"value,omitempty"`
}

func (o ScrumCustomField) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScrumCustomField struct{}"
	}

	return strings.Join([]string{"ScrumCustomField", string(data)}, " ")
}
