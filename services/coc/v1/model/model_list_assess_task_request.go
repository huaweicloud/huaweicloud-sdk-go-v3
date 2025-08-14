package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAssessTaskRequest Request Object
type ListAssessTaskRequest struct {

	// 分页参数
	Offset int32 `json:"offset"`

	// 每页显示的条目数量
	Limit int32 `json:"limit"`

	// 应用id
	ApplicationId *string `json:"application_id,omitempty"`

	// 评估任务状态
	AssessStatusList *[]string `json:"assess_status_list,omitempty"`
}

func (o ListAssessTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssessTaskRequest struct{}"
	}

	return strings.Join([]string{"ListAssessTaskRequest", string(data)}, " ")
}
