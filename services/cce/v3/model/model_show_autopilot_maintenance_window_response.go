package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutopilotMaintenanceWindowResponse Response Object
type ShowAutopilotMaintenanceWindowResponse struct {

	// 升级周期，从\"Monday,Tuesday,Wednesday,Thursday,Friday,Saturday,Sunday\"中进行选择，以英文逗号分隔 如：\"Friday,Saturday\"
	WeeklyPeriod *string `json:"weekly_period,omitempty"`

	// 升级起始时间（UTC时间）
	StartTime *string `json:"start_time,omitempty"`

	// 升级时长
	DurationHours *int32 `json:"duration_hours,omitempty"`

	// 升级版本范围，当前只支持小版本自动升级
	UpgradeScope   *string `json:"upgrade_scope,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAutopilotMaintenanceWindowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutopilotMaintenanceWindowResponse struct{}"
	}

	return strings.Join([]string{"ShowAutopilotMaintenanceWindowResponse", string(data)}, " ")
}
