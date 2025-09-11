package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNotificationMaskResourcesRequest Request Object
type ListNotificationMaskResourcesRequest struct {

	// **参数解释**： 屏蔽规则ID。    **约束限制**： 不涉及。 **取值范围**： 只能包含字母、数字，长度为[1,64]个字符。           **默认取值**： 不涉及。
	NotificationMaskId string `json:"notification_mask_id"`

	// 分页偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 分页大小
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListNotificationMaskResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotificationMaskResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListNotificationMaskResourcesRequest", string(data)}, " ")
}
