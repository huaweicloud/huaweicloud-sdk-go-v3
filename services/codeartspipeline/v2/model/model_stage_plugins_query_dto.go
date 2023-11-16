package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StagePluginsQueryDto struct {

	// 用于区分插件为流水线可使用/模板可使用
	UseCondition *string `json:"use_condition,omitempty"`

	// 微服务ID
	CompId *string `json:"comp_id,omitempty"`

	// 微服务名
	CompName *string `json:"comp_name,omitempty"`

	// 局点ID
	CloudId *string `json:"cloud_id,omitempty"`

	// 策略ID
	StrategyId *string `json:"strategy_id,omitempty"`

	// 流水线类型
	Category *string `json:"category,omitempty"`

	// 是否发布流水线
	PublishTab *string `json:"publish_tab,omitempty"`

	// 部署平台
	Platform *string `json:"platform,omitempty"`

	// 组件类型
	CompExtendType *string `json:"comp_extend_type,omitempty"`

	// 部署类型
	DeployType *string `json:"deploy_type,omitempty"`
}

func (o StagePluginsQueryDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StagePluginsQueryDto struct{}"
	}

	return strings.Join([]string{"StagePluginsQueryDto", string(data)}, " ")
}
