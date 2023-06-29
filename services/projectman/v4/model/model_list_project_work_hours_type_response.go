package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectWorkHoursTypeResponse Response Object
type ListProjectWorkHoursTypeResponse struct {

	// 工时类型列表
	WorkHoursTypes *[]WorkHoursType `json:"work_hours_types,omitempty"`

	// 总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListProjectWorkHoursTypeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectWorkHoursTypeResponse struct{}"
	}

	return strings.Join([]string{"ListProjectWorkHoursTypeResponse", string(data)}, " ")
}
