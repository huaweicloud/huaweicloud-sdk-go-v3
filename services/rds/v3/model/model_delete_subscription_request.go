package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSubscriptionRequest Request Object
type DeleteSubscriptionRequest struct {

	// 语言。默认en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *DeleteSubscriptionRequestBody `json:"body,omitempty"`
}

func (o DeleteSubscriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSubscriptionRequest struct{}"
	}

	return strings.Join([]string{"DeleteSubscriptionRequest", string(data)}, " ")
}
