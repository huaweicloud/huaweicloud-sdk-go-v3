package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TypeProperties struct {
	// Kill chain事件分类，仅当business为attack有效

	Killchain *string `json:"killchain,omitempty"`
	// Mitre Array 事件分类，仅当business为attack有效

	Ttps *string `json:"ttps,omitempty"`
	// 影响，适用全部类型

	Effects *string `json:"effects,omitempty"`
}

func (o TypeProperties) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TypeProperties struct{}"
	}

	return strings.Join([]string{"TypeProperties", string(data)}, " ")
}
