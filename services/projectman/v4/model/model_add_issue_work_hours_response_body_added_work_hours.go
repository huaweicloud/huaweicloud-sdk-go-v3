package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddIssueWorkHoursResponseBodyAddedWorkHours struct {

	// 工时id
	WorkHoursId *string `json:"work_hours_id,omitempty"`

	// 工作项id
	IssueId *int32 `json:"issue_id,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`

	// 用户昵称
	UserNickName *string `json:"user_nick_name,omitempty"`

	// 工时日期
	WorkDate *string `json:"work_date,omitempty"`

	// 工时数
	WorkHours *string `json:"work_hours,omitempty"`

	// 工时类型名称
	WorkHoursTypeName *string `json:"work_hours_type_name,omitempty"`
}

func (o AddIssueWorkHoursResponseBodyAddedWorkHours) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddIssueWorkHoursResponseBodyAddedWorkHours struct{}"
	}

	return strings.Join([]string{"AddIssueWorkHoursResponseBodyAddedWorkHours", string(data)}, " ")
}
