package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubscriptionTasksRequest Request Object
type ListSubscriptionTasksRequest struct {

	// **参数解释：** 每页显示的条目数量 **约束限制：** 不涉及 **取值范围：** 0-100 **默认取值：** 5 > 传入空或0时，会按默认处理
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 偏移量 > 表示从此偏移量开始查询  **约束限制：** 不涉及 **取值范围：** offset大于等于0 **默认取值：** 0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListSubscriptionTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubscriptionTasksRequest struct{}"
	}

	return strings.Join([]string{"ListSubscriptionTasksRequest", string(data)}, " ")
}
