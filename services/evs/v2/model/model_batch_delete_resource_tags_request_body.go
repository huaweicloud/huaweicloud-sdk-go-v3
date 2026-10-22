package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteResourceTagsRequestBody 删除标签的请求body体。
type BatchDeleteResourceTagsRequestBody struct {

	// **参数解释** 标签列表。 **约束限制** 不涉及。 **取值范围** 不涉及。 **默认取值** 不涉及。
	Tags []DeleteResourceTag `json:"tags"`
}

func (o BatchDeleteResourceTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteResourceTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteResourceTagsRequestBody", string(data)}, " ")
}
