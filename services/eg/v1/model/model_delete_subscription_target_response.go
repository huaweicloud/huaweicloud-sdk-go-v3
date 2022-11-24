package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
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
