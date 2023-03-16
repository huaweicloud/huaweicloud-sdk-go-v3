package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateVideoObjectMaskingTaskRequest struct {
	Body *CreateVideoObjectMaskingTaskRequestBody `json:"body,omitempty"`
}

func (o CreateVideoObjectMaskingTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVideoObjectMaskingTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateVideoObjectMaskingTaskRequest", string(data)}, " ")
}
