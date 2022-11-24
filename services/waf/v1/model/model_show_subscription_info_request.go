package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowSubscriptionInfoRequest struct {
}

func (o ShowSubscriptionInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSubscriptionInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowSubscriptionInfoRequest", string(data)}, " ")
}
