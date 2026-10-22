package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// V2BatchDeleteTagsRequestBody **参数解释**：批量删除的资源标签。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
type V2BatchDeleteTagsRequestBody struct {

	// **参数解释**：批量删除请求体标签列表，每个元素为一个V2DeleteResourceTag对象，包含key（标签键）和value（标签值）字段。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Tags []V2DeleteResourceTag `json:"tags"`

	// **参数解释**：批量删除请求体系统标签列表，每个元素为一个SystemTag对象，包含key（系统标签键）和value（系统标签值）字段。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	SysTags *[]SystemTag `json:"sys_tags,omitempty"`
}

func (o V2BatchDeleteTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "V2BatchDeleteTagsRequestBody struct{}"
	}

	return strings.Join([]string{"V2BatchDeleteTagsRequestBody", string(data)}, " ")
}
