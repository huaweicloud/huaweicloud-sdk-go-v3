package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteNotificationMasksResponse Response Object
type BatchDeleteNotificationMasksResponse struct {

	// 删除成功的屏蔽规则ID
	NotificationMaskIds *[]string `json:"notification_mask_ids,omitempty"`
	HttpStatusCode      int       `json:"-"`
}

func (o BatchDeleteNotificationMasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteNotificationMasksResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteNotificationMasksResponse", string(data)}, " ")
}
