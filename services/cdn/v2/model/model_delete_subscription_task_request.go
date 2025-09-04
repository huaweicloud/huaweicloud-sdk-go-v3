package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSubscriptionTaskRequest Request Object
type DeleteSubscriptionTaskRequest struct {

	// 订阅任务id
	Id int64 `json:"id"`
}

func (o DeleteSubscriptionTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSubscriptionTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteSubscriptionTaskRequest", string(data)}, " ")
}
