package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateTextToImageTaskRequest struct {
	Body *CreateTextToImageTaskRequestBody `json:"body,omitempty"`
}

func (o CreateTextToImageTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTextToImageTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateTextToImageTaskRequest", string(data)}, " ")
}
