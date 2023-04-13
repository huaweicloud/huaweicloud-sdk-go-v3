package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateImageTranslateTaskRequest struct {
	Body *CreateImageTranslateRequestBody `json:"body,omitempty"`
}

func (o CreateImageTranslateTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImageTranslateTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateImageTranslateTaskRequest", string(data)}, " ")
}
