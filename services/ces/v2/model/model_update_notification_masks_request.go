package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNotificationMasksRequest Request Object
type UpdateNotificationMasksRequest struct {

	// 屏蔽规则ID
	NotificationMaskId string `json:"notification_mask_id"`

	Body *UpdateNotificationMasksRequestBody `json:"body,omitempty"`
}

func (o UpdateNotificationMasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNotificationMasksRequest struct{}"
	}

	return strings.Join([]string{"UpdateNotificationMasksRequest", string(data)}, " ")
}
