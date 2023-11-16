package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SimpleOpenSourceRuleSetVo 开源治理策略列表详情
type SimpleOpenSourceRuleSetVo struct {

	// 开源治理策略ID
	Id *string `json:"id,omitempty"`

	// 开源治理策略名称
	Name *string `json:"name,omitempty"`

	// 开源治理策略级别（tenant-租户级，project-项目级）
	Level *string `json:"level,omitempty"`

	// 是否可用
	IsValid *bool `json:"is_valid,omitempty"`

	// 版本
	Version *string `json:"version,omitempty"`

	// 操作人
	Operator *string `json:"operator,omitempty"`

	// 是否系统策略
	IsPublic *bool `json:"is_public,omitempty"`

	// 是否老版
	IsLegacy *bool `json:"is_legacy,omitempty"`

	// 操作时间
	OperateTime *int64 `json:"operate_time,omitempty"`
}

func (o SimpleOpenSourceRuleSetVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SimpleOpenSourceRuleSetVo struct{}"
	}

	return strings.Join([]string{"SimpleOpenSourceRuleSetVo", string(data)}, " ")
}
