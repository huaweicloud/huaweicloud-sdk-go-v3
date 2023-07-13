package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNotificationRequest Request Object
type UpdateNotificationRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 订阅管理ID
	NotificationId int64 `json:"notification_id"`

	Body *UpdateNotificationRequestBody `json:"body,omitempty"`
}

func (o UpdateNotificationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNotificationRequest struct{}"
	}

	return strings.Join([]string{"UpdateNotificationRequest", string(data)}, " ")
}
