package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProjectStrategyResponse Response Object
type ShowProjectStrategyResponse struct {

	// 规则模版实例ID
	Id *string `json:"id,omitempty"`

	// 规则模版实例名称
	Name *string `json:"name,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`

	// 版本
	Version *string `json:"version,omitempty"`

	// 创建人
	Creator *string `json:"creator,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 最近更新人
	Updater *string `json:"updater,omitempty"`

	// 最近更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 是否生效
	IsValid *bool `json:"is_valid,omitempty"`

	// 规则实例集合
	RuleInstances  *[]RuleInstance `json:"rule_instances,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowProjectStrategyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectStrategyResponse struct{}"
	}

	return strings.Join([]string{"ShowProjectStrategyResponse", string(data)}, " ")
}
