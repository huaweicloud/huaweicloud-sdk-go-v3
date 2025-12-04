package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceConsumerGroupsResponse Response Object
type ListInstanceConsumerGroupsResponse struct {

	// **参数解释**： 所有消费组的信息。
	Groups *[]GroupInfoSimple `json:"groups,omitempty"`

	// **参数解释**： 消费组总数。 **取值范围**： 不涉及。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstanceConsumerGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceConsumerGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceConsumerGroupsResponse", string(data)}, " ")
}
