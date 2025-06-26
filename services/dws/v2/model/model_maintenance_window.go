package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MaintenanceWindow **参数解释**： 可维护时间段。 **取值范围**： 不涉及。
type MaintenanceWindow struct {

	// **参数解释**： 日期，范围：Mon、Tue、Wed、Thu、Fri、Sat、Sun。 **取值范围**： 不涉及。
	Day string `json:"day"`

	// **参数解释**： 开始时间，UTC时间，格式为HH:mm，例如：22:00。 时间必须是整点，且开始时间和结束时间必须间隔4小时。 **取值范围**： 不涉及。
	StartTime string `json:"start_time"`

	// **参数解释**： 结束时间，UTC时间，格式为HH:mm，例如：22:00。 时间必须是整点，且开始时间和结束时间必须间隔4小时。 **取值范围**： 不涉及。
	EndTime string `json:"end_time"`
}

func (o MaintenanceWindow) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MaintenanceWindow struct{}"
	}

	return strings.Join([]string{"MaintenanceWindow", string(data)}, " ")
}
