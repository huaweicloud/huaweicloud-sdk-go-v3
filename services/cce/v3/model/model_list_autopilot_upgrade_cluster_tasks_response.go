package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAutopilotUpgradeClusterTasksResponse Response Object
type ListAutopilotUpgradeClusterTasksResponse struct {

	// **参数解释：** API版本，默认为v3 **约束限制：** 不涉及 **取值范围：** - v3  **默认取值：** v3
	ApiVersion *string `json:"apiVersion,omitempty"`

	// **参数解释：** 资源类型 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Kind *string `json:"kind,omitempty"`

	Metadata *UpgradeTaskMetadata `json:"metadata,omitempty"`

	// **参数解释：** 集群升级任务列表 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Items          *[]UpgradeTaskResponseBody `json:"items,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListAutopilotUpgradeClusterTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAutopilotUpgradeClusterTasksResponse struct{}"
	}

	return strings.Join([]string{"ListAutopilotUpgradeClusterTasksResponse", string(data)}, " ")
}
