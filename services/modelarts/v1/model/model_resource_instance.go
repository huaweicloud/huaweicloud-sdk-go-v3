package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceInstance 训练作业资源信息及标签。
type ResourceInstance struct {

	// **参数解释**：资源详情，当前为空对象。 **取值范围**：不涉及。
	ResourceDetail *interface{} `json:"resource_detail,omitempty"`

	// **参数解释**：工作空间ID。 **取值范围**：不涉及。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释**：训练作业ID。 **取值范围**：不涉及。
	ResourceId *string `json:"resource_id,omitempty"`

	// **参数解释**：训练作业名称。 **取值范围**：不涉及。
	ResourceName *string `json:"resource_name,omitempty"`

	// **参数解释**：该训练作业上的标签列表。无标签的作业返回空数组。 **取值范围**：不涉及。
	Tags *[]TmsTagResp `json:"tags,omitempty"`
}

func (o ResourceInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceInstance struct{}"
	}

	return strings.Join([]string{"ResourceInstance", string(data)}, " ")
}
