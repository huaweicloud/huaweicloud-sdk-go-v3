package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NotificationResp struct {

	// **参数解释**： 告警通知类型。 **取值范围**： - notification：通知组或主题订阅。 - autoscaling：AS通知，仅在AS使用。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 告警状态发生变化时，被通知对象的列表。
	NotificationList *[]string `json:"notificationList,omitempty"`
}

func (o NotificationResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotificationResp struct{}"
	}

	return strings.Join([]string{"NotificationResp", string(data)}, " ")
}
