package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowProjectWorkHoursResponse struct {
	// 工时列表

	WorkHours *[]ShowProjectWorkHoursResponseBodyWorkHours `json:"work_hours,omitempty"`
	// 总数

	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowProjectWorkHoursResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectWorkHoursResponse struct{}"
	}

	return strings.Join([]string{"ShowProjectWorkHoursResponse", string(data)}, " ")
}
