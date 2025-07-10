package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRuleRestrictionResponse Response Object
type ShowRuleRestrictionResponse struct {

	// 应用管控开关，false：关闭，true：打开
	AppRestrictRuleSwitch *bool `json:"app_restrict_rule_switch,omitempty"`

	// 应用管控方式，0：禁止列表应用程序运行
	AppControlMode *int32 `json:"app_control_mode,omitempty"`

	// 周期性监控开关，false：关闭，true：打开
	AppPeriodicSwitch *bool `json:"app_periodic_switch,omitempty"`

	// 周期性监控间隔时间，单位分钟
	AppPeriodicInterval *int32 `json:"app_periodic_interval,omitempty"`

	// 强制kill应用开关，false：关闭，true：打开
	AppForceKillProcSwitch *bool `json:"app_force_kill_proc_switch,omitempty"`
	HttpStatusCode         int   `json:"-"`
}

func (o ShowRuleRestrictionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRuleRestrictionResponse struct{}"
	}

	return strings.Join([]string{"ShowRuleRestrictionResponse", string(data)}, " ")
}
