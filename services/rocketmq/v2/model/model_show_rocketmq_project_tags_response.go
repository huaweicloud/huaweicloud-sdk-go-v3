package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRocketmqProjectTagsResponse Response Object
type ShowRocketmqProjectTagsResponse struct {

	// **参数解释**： 总数。 **取值范围**： 不涉及。
	Total float32 `json:"total,omitempty"`

	// **参数解释**： 下个分页的offset。 **取值范围**： 不涉及。
	NextOffset *int32 `json:"next_offset,omitempty"`

	// **参数解释**： 上个分页的offset。 **取值范围**： 不涉及。
	PreviousOffset *int32 `json:"previous_offset,omitempty"`

	// **参数解释**： 标签列表。
	Tags           *[]TagMultyValueEntity `json:"tags,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowRocketmqProjectTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRocketmqProjectTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowRocketmqProjectTagsResponse", string(data)}, " ")
}
