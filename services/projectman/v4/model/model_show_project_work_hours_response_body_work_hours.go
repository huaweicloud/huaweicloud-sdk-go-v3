package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowProjectWorkHoursResponseBodyWorkHours struct {

	// 项目名称
	ProjectName *string `json:"project_name,omitempty" xml:"project_name"`

	// 用户昵称
	NickName *string `json:"nick_name,omitempty" xml:"nick_name"`

	// 用户id
	UserId *string `json:"user_id,omitempty" xml:"user_id"`

	// 用户名
	UserName *string `json:"user_name,omitempty" xml:"user_name"`

	// 工时日期
	WorkDate *string `json:"work_date,omitempty" xml:"work_date"`

	// 工时花费
	WorkHoursNum *string `json:"work_hours_num,omitempty" xml:"work_hours_num"`

	// 工时内容
	Summary *string `json:"summary,omitempty" xml:"summary"`

	// 工时类型
	WorkHoursTypeName *string `json:"work_hours_type_name,omitempty" xml:"work_hours_type_name"`

	// 工作项id
	IssueId *int32 `json:"issue_id,omitempty" xml:"issue_id"`

	// 工作项类型
	IssueType *string `json:"issue_type,omitempty" xml:"issue_type"`

	// 工作项标题
	Subject *string `json:"subject,omitempty" xml:"subject"`

	// 工作项创建时间
	CreatedTime *string `json:"created_time,omitempty" xml:"created_time"`

	// 工作项结束时间
	ClosedTime *string `json:"closed_time,omitempty" xml:"closed_time"`
}

func (o ShowProjectWorkHoursResponseBodyWorkHours) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectWorkHoursResponseBodyWorkHours struct{}"
	}

	return strings.Join([]string{"ShowProjectWorkHoursResponseBodyWorkHours", string(data)}, " ")
}
