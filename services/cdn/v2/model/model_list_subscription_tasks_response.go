package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubscriptionTasksResponse Response Object
type ListSubscriptionTasksResponse struct {

	// **参数解释：** 订阅任务总数 **取值范围：** 不涉及
	Total *int32 `json:"total,omitempty"`

	Data           *[]SubscriptionTask `json:"data,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListSubscriptionTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubscriptionTasksResponse struct{}"
	}

	return strings.Join([]string{"ListSubscriptionTasksResponse", string(data)}, " ")
}
