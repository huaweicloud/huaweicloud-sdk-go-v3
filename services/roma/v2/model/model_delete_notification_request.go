package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteNotificationRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 订阅管理ID

	NotificationId int64 `json:"notification_id"`
}

func (o DeleteNotificationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNotificationRequest struct{}"
	}

	return strings.Join([]string{"DeleteNotificationRequest", string(data)}, " ")
}
