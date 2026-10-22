package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteCustomLlmRequest Request Object
type BatchDeleteCustomLlmRequest struct {
	Body *CustomLlmBatchDeleteBody `json:"body,omitempty"`
}

func (o BatchDeleteCustomLlmRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteCustomLlmRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteCustomLlmRequest", string(data)}, " ")
}
