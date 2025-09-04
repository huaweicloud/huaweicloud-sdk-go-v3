package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSubscriptionTaskResponse Response Object
type UpdateSubscriptionTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateSubscriptionTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubscriptionTaskResponse struct{}"
	}

	return strings.Join([]string{"UpdateSubscriptionTaskResponse", string(data)}, " ")
}
