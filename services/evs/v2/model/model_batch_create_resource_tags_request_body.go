package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateResourceTagsRequestBody 创建资源标签的请求body体。
type BatchCreateResourceTagsRequestBody struct {

	// **参数解释** 标签列表。 **约束限制** 不涉及。 **取值范围** 不涉及。 **默认取值** 不涉及。
	Tags []CreateResourceTag `json:"tags"`
}

func (o BatchCreateResourceTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateResourceTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateResourceTagsRequestBody", string(data)}, " ")
}
