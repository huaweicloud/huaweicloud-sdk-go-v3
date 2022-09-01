package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type OperateSubscriptionRequest struct {
	Body *SubscriptionOperateReq `json:"body,omitempty" xml:"body"`
}

func (o OperateSubscriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateSubscriptionRequest struct{}"
	}

	return strings.Join([]string{"OperateSubscriptionRequest", string(data)}, " ")
}
