package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSubscriptionUserRequest Request Object
type CreateSubscriptionUserRequest struct {
	Body *CreateSubscriptionUserRequestBody `json:"body,omitempty"`
}

func (o CreateSubscriptionUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubscriptionUserRequest struct{}"
	}

	return strings.Join([]string{"CreateSubscriptionUserRequest", string(data)}, " ")
}
