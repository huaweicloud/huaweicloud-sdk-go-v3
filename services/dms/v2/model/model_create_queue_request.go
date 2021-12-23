package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateQueueRequest struct {
	Body *CreateQueueReq `json:"body,omitempty"`
}

func (o CreateQueueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateQueueRequest struct{}"
	}

	return strings.Join([]string{"CreateQueueRequest", string(data)}, " ")
}
