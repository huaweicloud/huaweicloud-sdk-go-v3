package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSubscriptionTaskResponse Response Object
type CreateSubscriptionTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateSubscriptionTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubscriptionTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateSubscriptionTaskResponse", string(data)}, " ")
}
