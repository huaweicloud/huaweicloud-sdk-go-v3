package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAssessTaskRequest Request Object
type CreateAssessTaskRequest struct {
	Body *CreateAssessTaskRequestBody `json:"body,omitempty"`
}

func (o CreateAssessTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAssessTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateAssessTaskRequest", string(data)}, " ")
}
