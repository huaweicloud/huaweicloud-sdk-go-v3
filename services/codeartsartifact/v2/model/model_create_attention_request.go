package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAttentionRequest Request Object
type CreateAttentionRequest struct {
	Body *AttentionDo `json:"body,omitempty"`
}

func (o CreateAttentionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAttentionRequest struct{}"
	}

	return strings.Join([]string{"CreateAttentionRequest", string(data)}, " ")
}
