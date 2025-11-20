package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RefreshSubscriptionResponse Response Object
type RefreshSubscriptionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RefreshSubscriptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RefreshSubscriptionResponse struct{}"
	}

	return strings.Join([]string{"RefreshSubscriptionResponse", string(data)}, " ")
}
