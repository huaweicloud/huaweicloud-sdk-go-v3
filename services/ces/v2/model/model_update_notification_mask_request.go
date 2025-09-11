package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNotificationMaskRequest Request Object
type UpdateNotificationMaskRequest struct {

	// **参数解释**： 屏蔽规则ID。    **约束限制**： 不涉及。 **取值范围**： 只能包含字母、数字，长度为[1,64]个字符。           **默认取值**： 不涉及。
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
