package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAutopilotPreCheckTasksResponse Response Object
type ListAutopilotPreCheckTasksResponse struct {

	// **参数解释：** API版本，默认为v3 **约束限制：** 不涉及 **取值范围：** - v3  **默认取值：** v3
	ApiVersion *string `json:"apiVersion,omitempty"`

	// **参数解释：** 类型 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Kind *string `json:"kind,omitempty"`

	Metadata *Metadata `json:"metadata,omitempty"`

	// **参数解释：** 集群检查任务列表 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Items          *[]PrecheckClusterTask `json:"items,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListAutopilotPreCheckTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAutopilotPreCheckTasksResponse struct{}"
	}

	return strings.Join([]string{"ListAutopilotPreCheckTasksResponse", string(data)}, " ")
}
