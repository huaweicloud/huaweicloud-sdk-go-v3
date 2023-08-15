package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScalingNotificationsResponse Response Object
type ListScalingNotificationsResponse struct {

	// 伸缩组通知列表。
	Topics         *[]Topics `json:"topics,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListScalingNotificationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScalingNotificationsResponse struct{}"
	}

	return strings.Join([]string{"ListScalingNotificationsResponse", string(data)}, " ")
}
