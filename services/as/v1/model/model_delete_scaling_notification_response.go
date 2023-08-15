package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteScalingNotificationResponse Response Object
type DeleteScalingNotificationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteScalingNotificationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScalingNotificationResponse struct{}"
	}

	return strings.Join([]string{"DeleteScalingNotificationResponse", string(data)}, " ")
}
