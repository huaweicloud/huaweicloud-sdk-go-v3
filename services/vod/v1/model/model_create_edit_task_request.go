package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEditTaskRequest Request Object
type CreateEditTaskRequest struct {
	Body *CreateEditTaskRequestBody `json:"body,omitempty"`
}

func (o CreateEditTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEditTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateEditTaskRequest", string(data)}, " ")
}
