package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateNotificationMasksResponse Response Object
type BatchUpdateNotificationMasksResponse struct {

	// 创建成功的关联ID列表
	RelationIds *[]string `json:"relation_ids,omitempty"`

	// 屏蔽规则ID
	NotificationMaskId *string `json:"notification_mask_id,omitempty"`
	HttpStatusCode     int     `json:"-"`
}

func (o BatchUpdateNotificationMasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateNotificationMasksResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateNotificationMasksResponse", string(data)}, " ")
}
