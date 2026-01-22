package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyPullTaskRequest Request Object
type ModifyPullTaskRequest struct {
	Body *ModifyLivePullStreamTask `json:"body,omitempty"`
}

func (o ModifyPullTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyPullTaskRequest struct{}"
	}

	return strings.Join([]string{"ModifyPullTaskRequest", string(data)}, " ")
}
