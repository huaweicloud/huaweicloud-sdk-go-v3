package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTimerRuleDetails 修改启停规则规格信息
type UpdateTimerRuleDetails struct {

	// 定时启停规则名称。
	Name string `json:"name"`

	// 定时启停规则的类型：stop/start。
	Type string `json:"type"`

	// 定时启停规则状态（是否开启）：on/off。
	Status string `json:"status"`

	// 定时启停规则所包含的所有应用，只在生效范围为application的时候需要填写。
	Apps *[]AppInfo `json:"apps,omitempty"`

	// 在定时启停规则所包含的所有组件，只在生效范围为component的时候需要填写。
	Components *[]ComponentInfo `json:"components,omitempty"`

	// cron表达式。
	Cron string `json:"cron"`

	// 定时启停规则生效范围: component/application/environment。
	EffectiveRange string `json:"effective_range"`

	// 定时启停规则的定时类别: onetime/periodic。
	EffectivePolicy string `json:"effective_policy"`
}

func (o UpdateTimerRuleDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTimerRuleDetails struct{}"
	}

	return strings.Join([]string{"UpdateTimerRuleDetails", string(data)}, " ")
}
