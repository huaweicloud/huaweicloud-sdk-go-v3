package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPoolsRequest Request Object
type ListPoolsRequest struct {

	// **参数解释**：工作空间ID。[获取方法请参见[查询工作空间列表](ListWorkspace.xml)。](tag:hc,hk) 未创建工作空间时默认值为“0”，存在创建并使用的工作空间，以实际取值为准。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：0。
	WorkspaceId *string `json:"workspaceId,omitempty"`

	// **参数解释**：资源池标签筛选。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	LabelSelector *string `json:"labelSelector,omitempty"`

	// **参数解释**：资源池状态。 **约束限制**：不涉及。 **取值范围**：可选值如下： - created: 创建成功的资源池。 - failed：创建失败的资源池，创建失败的资源池记录保留3天。 - creating：创建中的资源池 **默认取值**：不涉及。
	Status *string `json:"status,omitempty"`
}

func (o ListPoolsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPoolsRequest struct{}"
	}

	return strings.Join([]string{"ListPoolsRequest", string(data)}, " ")
}
