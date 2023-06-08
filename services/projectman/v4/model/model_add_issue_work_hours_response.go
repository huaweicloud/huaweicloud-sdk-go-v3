package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AddIssueWorkHoursResponse struct {

	// 已添加的工时列表
	AddedWorkHours *[]AddIssueWorkHoursResponseBodyAddedWorkHours `json:"added_work_hours,omitempty"`
	HttpStatusCode int                                            `json:"-"`
}

func (o AddIssueWorkHoursResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddIssueWorkHoursResponse struct{}"
	}

	return strings.Join([]string{"AddIssueWorkHoursResponse", string(data)}, " ")
}
