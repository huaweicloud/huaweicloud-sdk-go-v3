package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TimerRuleDetails 启停规则列表。
type TimerRuleDetails struct {

	// 定时启停规则ID，在创建定时启停规则时会忽略。
	Id *string `json:"id,omitempty"`

	// 定时启停规则名称。
	Name *string `json:"name,omitempty"`

	// 定时启停规则的类型：stop/start。
	Type *string `json:"type,omitempty"`

	// 定时启停规则状态（是否开启）：on/off。
	Status *string `json:"status,omitempty"`

	// 环境ID。
	EnvId *string `json:"env_id,omitempty"`

	// 定时启停规则所包含的所有应用，只在生效范围为application的时候需要填写。
	Apps *[]AppInfo `json:"apps,omitempty"`

	// 在定时启停规则所包含的所有组件，只在生效范围为component的时候需要填写。
	Components *[]ComponentInfo `json:"components,omitempty"`

	// 定时启停规则包含的组件个数，在创建定时启停规则时会忽略。
	ComponentNumber *int32 `json:"component_number,omitempty"`

	// cron表达式。
	Cron *string `json:"cron,omitempty"`

	// 定时启停规则生效范围: component/application/environment。
	EffectiveRange *string `json:"effective_range,omitempty"`

	// 定时启停规则的定时类别: onetime/periodic。
	EffectivePolicy *string `json:"effective_policy,omitempty"`

	// 上次执行的状态：abnormal/normal/executing，在创建定时启停规则时会忽略。
	LastExecutionStatus *string `json:"last_execution_status,omitempty"`
}

func (o TimerRuleDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TimerRuleDetails struct{}"
	}

	return strings.Join([]string{"TimerRuleDetails", string(data)}, " ")
}
