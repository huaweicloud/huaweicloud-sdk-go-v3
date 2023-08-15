package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTopicRequest Request Object
type CreateTopicRequest struct {
	Body *CreateTopicRequestBody `json:"body,omitempty"`
}

func (o CreateTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTopicRequest struct{}"
	}

	return strings.Join([]string{"CreateTopicRequest", string(data)}, " ")
}
