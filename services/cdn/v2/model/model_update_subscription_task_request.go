package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSubscriptionTaskRequest Request Object
type UpdateSubscriptionTaskRequest struct {

	// 订阅任务id
	Id int64 `json:"id"`

	Body *SubscriptionTaskVo `json:"body,omitempty"`
}

func (o UpdateSubscriptionTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubscriptionTaskRequest struct{}"
	}

	return strings.Join([]string{"UpdateSubscriptionTaskRequest", string(data)}, " ")
}
