package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MaintenanceWindow 可维护时间段。
type MaintenanceWindow struct {

	// 日期，范围：Mon、Tue、Wed、Thu、Fri、Sat、Sun。
	Day string `json:"day"`

	// 开始时间，UTC时间，格式为HH:mm，例如：22:00。 - 时间必须是整点。 - 开始时间和结束时间必须间隔4小时。
	StartTime string `json:"start_time"`

	// 结束时间，UTC时间，格式为HH:mm，例如：02:00。 - 时间必须是整点。 - 开始时间和结束时间必须间隔4小时。
	EndTime string `json:"end_time"`
}

func (o MaintenanceWindow) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MaintenanceWindow struct{}"
	}

	return strings.Join([]string{"MaintenanceWindow", string(data)}, " ")
}
