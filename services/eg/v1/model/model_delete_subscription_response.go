package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteSubscriptionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSubscriptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSubscriptionResponse struct{}"
	}

	return strings.Join([]string{"DeleteSubscriptionResponse", string(data)}, " ")
}
