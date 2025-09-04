package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSubscriptionTaskRequest Request Object
type CreateSubscriptionTaskRequest struct {
	Body *SubscriptionTaskVo `json:"body,omitempty"`
}

func (o CreateSubscriptionTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubscriptionTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateSubscriptionTaskRequest", string(data)}, " ")
}
