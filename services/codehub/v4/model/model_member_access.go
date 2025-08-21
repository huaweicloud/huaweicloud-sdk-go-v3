package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MemberAccess struct {

	// **参数解释：** 访问级别。 **约束限制：** 不涉及。
	AccessLevel *int32 `json:"access_level,omitempty"`

	// **参数解释：** 通知级别。 **约束限制：** 不涉及。
	NotificationLevel *int32 `json:"notification_level,omitempty"`
}

func (o MemberAccess) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MemberAccess struct{}"
	}

	return strings.Join([]string{"MemberAccess", string(data)}, " ")
}
