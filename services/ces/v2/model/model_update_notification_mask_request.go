package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNotificationMaskRequest Request Object
type UpdateNotificationMaskRequest struct {

	// 屏蔽规则ID
	NotificationMaskId string `json:"notification_mask_id"`

	Body *UpdateNotificationMasksRequestBody `json:"body,omitempty"`
}

func (o UpdateNotificationMaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNotificationMaskRequest struct{}"
	}

	return strings.Join([]string{"UpdateNotificationMaskRequest", string(data)}, " ")
}
