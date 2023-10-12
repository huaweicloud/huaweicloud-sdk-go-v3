package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteNotificationMasksRequestBody struct {

	// 屏蔽规则编号
	NotificationMaskIds []string `json:"notification_mask_ids"`
}

func (o BatchDeleteNotificationMasksRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteNotificationMasksRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteNotificationMasksRequestBody", string(data)}, " ")
}
