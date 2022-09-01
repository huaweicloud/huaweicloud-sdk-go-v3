package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListTaskSuccessRateRequest struct {

	// 项目id
	ProjectId string `json:"project_id" xml:"project_id"`

	Body *TasksSuccessRateQuery `json:"body,omitempty" xml:"body"`
}

func (o ListTaskSuccessRateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTaskSuccessRateRequest struct{}"
	}

	return strings.Join([]string{"ListTaskSuccessRateRequest", string(data)}, " ")
}
