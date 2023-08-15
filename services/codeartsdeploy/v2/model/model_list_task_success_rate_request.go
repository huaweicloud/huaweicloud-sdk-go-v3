package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTaskSuccessRateRequest Request Object
type ListTaskSuccessRateRequest struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	Body *TasksSuccessRateQuery `json:"body,omitempty"`
}

func (o ListTaskSuccessRateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTaskSuccessRateRequest struct{}"
	}

	return strings.Join([]string{"ListTaskSuccessRateRequest", string(data)}, " ")
}
