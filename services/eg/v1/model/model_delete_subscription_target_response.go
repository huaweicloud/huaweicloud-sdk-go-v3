package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSubscriptionTargetResponse Response Object
type DeleteSubscriptionTargetResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSubscriptionTargetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSubscriptionTargetResponse struct{}"
	}

	return strings.Join([]string{"DeleteSubscriptionTargetResponse", string(data)}, " ")
}
