package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAiAppRequest Request Object
type CreateAiAppRequest struct {
	Body *AiAppCreateRequest `json:"body,omitempty"`
}

func (o CreateAiAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAiAppRequest struct{}"
	}

	return strings.Join([]string{"CreateAiAppRequest", string(data)}, " ")
}
