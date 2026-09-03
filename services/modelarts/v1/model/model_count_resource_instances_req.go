package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountResourceInstancesReq **参数解释**：按标签统计训练作业资源数量的请求体。
type CountResourceInstancesReq struct {

	// **参数解释**：标签过滤条件，返回同时包含列表中所有标签的训练作业。 **约束限制**：标签个数不能超过系统允许的最大标签数。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Tags *[]CountResourceInstancesReqTags `json:"tags,omitempty"`

	// **参数解释**：模糊匹配条件，支持按资源名称等字段进行模糊查询。 **约束限制**：最多1个匹配条件。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Matches *[]CountResourceInstancesReqMatches `json:"matches,omitempty"`

	// **参数解释**：工作空间ID。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：0。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释**：是否查询不带任何标签的训练作业。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：false。
	WithoutAnyTag *bool `json:"without_any_tag,omitempty"`
}

func (o CountResourceInstancesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountResourceInstancesReq struct{}"
	}

	return strings.Join([]string{"CountResourceInstancesReq", string(data)}, " ")
}
