package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type Resource struct {
	// 配额类型。  枚举值说明：  alarm，告警规则

	Type string `json:"type"`
	// 已使用配额数。

	Used int32 `json:"used"`
	// 单位。

	Unit string `json:"unit"`
	// 配额总数。

	Quota int32 `json:"quota"`
}

func (o Resource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Resource struct{}"
	}

	return strings.Join([]string{"Resource", string(data)}, " ")
}
