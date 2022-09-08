package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListTaskSuccessRateResponse struct {

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 项目名称
	ProjectName *string `json:"project_name,omitempty"`

	// 任务执行开始时间范围的左边界（包含），格式yyyy-MM-dd
	StartTime *string `json:"start_time,omitempty"`

	// 任务执行开始时间范围的右边界（包含），格式yyyy-MM-dd 。最大时间范围为1年。
	EndTime *string `json:"end_time,omitempty"`

	// 任务的成功率列表
	TasksSuccessRate *[]TaskSuccessRate `json:"tasks_success_rate,omitempty"`
	HttpStatusCode   int                `json:"-"`
}

func (o ListTaskSuccessRateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTaskSuccessRateResponse struct{}"
	}

	return strings.Join([]string{"ListTaskSuccessRateResponse", string(data)}, " ")
}
