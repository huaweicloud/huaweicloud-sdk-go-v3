package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceQuotas
type ResourceQuotas struct {

	// 配额类型。  枚举值说明：  alarm，告警规则
	Type string `json:"type"`

	// 已使用配额数。
	Used int32 `json:"used"`

	// 单位。
	Unit string `json:"unit"`

	// 配额总数。
	Quota int32 `json:"quota"`
}

func (o ResourceQuotas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceQuotas struct{}"
	}

	return strings.Join([]string{"ResourceQuotas", string(data)}, " ")
}
