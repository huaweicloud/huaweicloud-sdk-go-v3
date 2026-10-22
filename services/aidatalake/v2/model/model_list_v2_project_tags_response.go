package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListV2ProjectTagsResponse Response Object
type ListV2ProjectTagsResponse struct {

	// **参数解释**：项目标签总记录数。 **取值范围**：0~2147483647。
	TotalCount *int32 `json:"total_count,omitempty"`

	// **参数解释**：项目标签列表，每个元素为一个V2ProjectResourceTag对象，包含key（标签键）和values（标签值列表）字段。
	Tags *[]V2ProjectResourceTag `json:"tags,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListV2ProjectTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListV2ProjectTagsResponse struct{}"
	}

	return strings.Join([]string{"ListV2ProjectTagsResponse", string(data)}, " ")
}
