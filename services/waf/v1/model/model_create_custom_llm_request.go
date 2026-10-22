package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCustomLlmRequest Request Object
type CreateCustomLlmRequest struct {
	Body *AiguardCustomLlmBody `json:"body,omitempty"`
}

func (o CreateCustomLlmRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCustomLlmRequest struct{}"
	}

	return strings.Join([]string{"CreateCustomLlmRequest", string(data)}, " ")
}
