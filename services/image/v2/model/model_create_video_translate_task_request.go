package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateVideoTranslateTaskRequest struct {
	Body *CreateVideoTranslateTaskRequestBody `json:"body,omitempty"`
}

func (o CreateVideoTranslateTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVideoTranslateTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateVideoTranslateTaskRequest", string(data)}, " ")
}
