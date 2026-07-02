package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRabbitMqProjectTagsResponse Response Object
type ShowRabbitMqProjectTagsResponse struct {

	// **参数解释**： 查询结果总数。 **取值范围**： 不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**： 下一个偏移量。 **取值范围**： 不涉及。
	NextOffset *int32 `json:"next_offset,omitempty"`

	// **参数解释**： 前一个偏移量。 **取值范围**： 不涉及。
	PreviousOffset *int32 `json:"previous_offset,omitempty"`

	// **参数解释**： 标签列表。
	Tags           *[]TagMultyValueEntity `json:"tags,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowRabbitMqProjectTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRabbitMqProjectTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowRabbitMqProjectTagsResponse", string(data)}, " ")
}
