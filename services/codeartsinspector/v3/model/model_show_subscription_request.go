package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSubscriptionRequest Request Object
type ShowSubscriptionRequest struct {

	// service
	Service string `json:"service"`
}

func (o ShowSubscriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSubscriptionRequest struct{}"
	}

	return strings.Join([]string{"ShowSubscriptionRequest", string(data)}, " ")
}
