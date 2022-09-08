package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type HandleNotificationRequest struct {
	Body *HandleNotificationRequestBody `json:"body,omitempty"`
}

func (o HandleNotificationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HandleNotificationRequest struct{}"
	}

	return strings.Join([]string{"HandleNotificationRequest", string(data)}, " ")
}
