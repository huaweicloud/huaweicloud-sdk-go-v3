package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEventSubRequest Request Object
type DeleteEventSubRequest struct {

	// **参数解释**： 事件订阅ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	EventSubId string `json:"event_sub_id"`
}

func (o DeleteEventSubRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEventSubRequest struct{}"
	}

	return strings.Join([]string{"DeleteEventSubRequest", string(data)}, " ")
}
