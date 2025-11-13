package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CollectSubscriptionMonitorRequest Request Object
type CollectSubscriptionMonitorRequest struct {

	// 语言。默认en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 订阅ID
	SubscriptionId string `json:"subscription_id"`
}

func (o CollectSubscriptionMonitorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectSubscriptionMonitorRequest struct{}"
	}

	return strings.Join([]string{"CollectSubscriptionMonitorRequest", string(data)}, " ")
}
