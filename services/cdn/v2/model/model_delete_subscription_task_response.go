package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSubscriptionTaskResponse Response Object
type DeleteSubscriptionTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSubscriptionTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSubscriptionTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteSubscriptionTaskResponse", string(data)}, " ")
}
