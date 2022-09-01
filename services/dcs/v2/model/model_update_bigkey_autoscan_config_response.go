package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateBigkeyAutoscanConfigResponse struct {

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id"`

	// 是否开启自动分析
	EnableAutoScan *bool `json:"enable_auto_scan,omitempty" xml:"enable_auto_scan"`

	// 每日分析时间，时间格式为21:00
	ScheduleAt *[]string `json:"schedule_at,omitempty" xml:"schedule_at"`

	// 配置更新时间，时间格式为2020-06-15T02:21:18.669Z
	UpdatedAt      *string `json:"updated_at,omitempty" xml:"updated_at"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateBigkeyAutoscanConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBigkeyAutoscanConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateBigkeyAutoscanConfigResponse", string(data)}, " ")
}
