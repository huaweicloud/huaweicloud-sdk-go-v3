package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceSubscriptionRequest Request Object
type ListInstanceSubscriptionRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o ListInstanceSubscriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceSubscriptionRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceSubscriptionRequest", string(data)}, " ")
}
