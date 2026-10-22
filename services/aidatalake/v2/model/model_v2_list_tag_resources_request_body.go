package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// V2ListTagResourcesRequestBody **参数解释**：标签管理界面查询资源实例列表的请求体。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
type V2ListTagResourcesRequestBody struct {

	// **参数解释**：不包含任意一个标签。 **约束限制**：不涉及。 **取值范围**：   - true：查询不包含任意一个标签的资源，此时忽略“tags”过滤字段。   - false：该条件不生效，即返回所有资源或按“tags”，“matches”等条件过滤。 **默认取值**：不涉及。
	WithoutAnyTag *bool `json:"without_any_tag,omitempty"`

	// **参数解释**：包含标签，最多包含20个key，每个key下面的value最多20个，每个key对应的value可以为空数组但结构体不能缺失。 **约束限制**：最大数量为20个。
	Tags *[]V2ResourceTagParam `json:"tags,omitempty"`

	// **参数解释**：前缀搜索字段。key为固定字典值，当前只支持resource_name，表示基于资源名称前缀搜索。value为匹配的值。 **约束限制**：不涉及。
	Matches *[]V2TagMatch `json:"matches,omitempty"`

	// **参数解释**：系统标签列表，目前只包含一个tag结构体。key下面只包含一个value。 **约束限制**：不涉及。
	SysTags *[]V2SystemTagParam `json:"sys_tags,omitempty"`
}

func (o V2ListTagResourcesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "V2ListTagResourcesRequestBody struct{}"
	}

	return strings.Join([]string{"V2ListTagResourcesRequestBody", string(data)}, " ")
}
