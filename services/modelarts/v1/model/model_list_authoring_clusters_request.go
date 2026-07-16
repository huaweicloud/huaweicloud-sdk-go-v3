package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuthoringClustersRequest Request Object
type ListAuthoringClustersRequest struct {

	// **参数解释**：资源池类型。 **约束限制**：不涉及。 **取值范围**：枚举类型，取值如下： - MANAGED： 公共池。 - DEDICATED：专属池。  **默认取值**：不涉及。
	Type string `json:"type"`

	// **参数解释**：工作空间ID。[获取方法请参见[查询工作空间列表](ListWorkspace.xml)。](tag:hc) **约束限制**：存在并使用的工作空间。 **取值范围**：不涉及。 **默认取值**：“0”。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释**：每一页显示实例的数量。 **约束限制**：不涉及。 **取值范围**：大于等于0。 **默认取值**：1000。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：数据条目偏移量。 **约束限制**：不涉及。 **取值范围**：大于等于0。 **默认取值**：0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：作业类型。 **约束限制**：不涉及。 **取值范围**：枚举类型，取值如下： - NOTEBOOK：开发环境 - TRAIN：训练作业 - INFER：推理作业  **默认取值**：NOTEBOOK。
	Scope *string `json:"scope,omitempty"`
}

func (o ListAuthoringClustersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuthoringClustersRequest struct{}"
	}

	return strings.Join([]string{"ListAuthoringClustersRequest", string(data)}, " ")
}
