package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// V2BatchCreateTagsRequestBody **参数解释**：批量打资源标签请求体。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
type V2BatchCreateTagsRequestBody struct {

	// **参数解释**：批量新增请求体标签列表，每个元素为一个V2ResourceTag对象，包含key（标签键）和value（标签值）字段。 **约束限制**：最多20个标签。 **默认取值**：不涉及。
	Tags *[]V2ResourceTag `json:"tags,omitempty"`

	// **参数解释**：系统标签列表，每个元素为一个SystemTag对象，包含key（系统标签键）和value（系统标签值）字段。 **约束限制**：不涉及。
	SysTags *[]SystemTag `json:"sys_tags,omitempty"`
}

func (o V2BatchCreateTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "V2BatchCreateTagsRequestBody struct{}"
	}

	return strings.Join([]string{"V2BatchCreateTagsRequestBody", string(data)}, " ")
}
