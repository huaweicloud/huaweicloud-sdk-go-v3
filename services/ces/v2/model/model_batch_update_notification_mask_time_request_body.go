package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateNotificationMaskTimeRequestBody 修改通知屏蔽时间请求体
type BatchUpdateNotificationMaskTimeRequestBody struct {

	// 关联编号
	NotificationMaskIds []string `json:"notification_mask_ids"`

	MaskType *MaskType `json:"mask_type"`

	// 屏蔽起始日期，yyyy-MM-dd。
	StartDate *string `json:"start_date,omitempty"`

	// 屏蔽起始时间，HH:mm:ss。
	StartTime *string `json:"start_time,omitempty"`

	// 屏蔽截止日期，yyyy-MM-dd。
	EndDate *string `json:"end_date,omitempty"`

	// 屏蔽截止时间，HH:mm:ss。
	EndTime *string `json:"end_time,omitempty"`

	// 时区，形如：\"GMT-08:00\"、\"GMT+08:00\"、\"GMT+0:00\"
	EffectiveTimezone *string `json:"effective_timezone,omitempty"`
}

func (o BatchUpdateNotificationMaskTimeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateNotificationMaskTimeRequestBody struct{}"
	}

	return strings.Join([]string{"BatchUpdateNotificationMaskTimeRequestBody", string(data)}, " ")
}
