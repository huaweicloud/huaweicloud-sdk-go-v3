package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSubscriptionTaskRequest Request Object
type UpdateSubscriptionTaskRequest struct {

	// **参数解释：** 订阅任务id **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Id int64 `json:"id"`

	Body *SubscriptionTaskVo `json:"body,omitempty"`
}

func (o UpdateSubscriptionTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubscriptionTaskRequest struct{}"
	}

	return strings.Join([]string{"UpdateSubscriptionTaskRequest", string(data)}, " ")
}
