package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type Resource struct {

	// 配额类型。  枚举值说明：  alarm，告警规则
	Type string `json:"type" xml:"type"`

	// 已使用配额数。
	Used int32 `json:"used" xml:"used"`

	// 单位。
	Unit string `json:"unit" xml:"unit"`

	// 配额总数。
	Quota int32 `json:"quota" xml:"quota"`
}

func (o Resource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Resource struct{}"
	}

	return strings.Join([]string{"Resource", string(data)}, " ")
}
