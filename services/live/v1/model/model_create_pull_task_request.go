package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePullTaskRequest Request Object
type CreatePullTaskRequest struct {
	Body *LivePullStreamTask `json:"body,omitempty"`
}

func (o CreatePullTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePullTaskRequest struct{}"
	}

	return strings.Join([]string{"CreatePullTaskRequest", string(data)}, " ")
}
