package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceInstancesFilterReq 查询训练作业资源列表的请求体。
type ResourceInstancesFilterReq struct {

	// **参数解释**：工作空间ID。未创建工作空间时默认值为\"0\"，存在创建并使用的工作空间，以实际取值为准。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：0。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释**：标签筛选条件，按标签key-value对筛选作业。同一key下多个value为OR关系，不同key之间为AND关系。 **约束限制**：同一key的values不能重复，不同key不能重复。 **取值范围**：不涉及。
	Tags *[]MutiValueTag `json:"tags,omitempty"`

	// **参数解释**：是否查询没有任何标签的作业。 **约束限制**：设为true时忽略tags筛选条件。 **取值范围**： - true：仅查询无标签的作业 - false：按tags条件筛选 **默认取值**：false。
	WithoutAnyTag *bool `json:"without_any_tag,omitempty"`

	// **参数解释**：资源名称搜索条件。 **约束限制**：最多支持1个匹配项，且key必须为resource_name。 **取值范围**：不涉及。
	Matches *[]Match `json:"matches,omitempty"`
}

func (o ResourceInstancesFilterReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceInstancesFilterReq struct{}"
	}

	return strings.Join([]string{"ResourceInstancesFilterReq", string(data)}, " ")
}
