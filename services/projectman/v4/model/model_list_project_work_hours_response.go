package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListProjectWorkHoursResponse struct {
	// 工时列表

	WorkHours *[]ShowProjectWorkHoursResponseBodyWorkHours `json:"work_hours,omitempty"`
	// 总数

	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListProjectWorkHoursResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectWorkHoursResponse struct{}"
	}

	return strings.Join([]string{"ListProjectWorkHoursResponse", string(data)}, " ")
}
