package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNotificationMasksResponse Response Object
type ListNotificationMasksResponse struct {

	// **参数解释**： 通知屏蔽列表
	NotificationMasks *[]ListNotificationMaskRespNotificationMasks `json:"notification_masks,omitempty"`

	// **参数解释**： 通知屏蔽列表总数 **取值范围**： [0,99999]
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListNotificationMasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotificationMasksResponse struct{}"
	}

	return strings.Join([]string{"ListNotificationMasksResponse", string(data)}, " ")
}
