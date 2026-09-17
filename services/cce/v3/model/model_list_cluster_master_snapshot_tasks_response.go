package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterMasterSnapshotTasksResponse Response Object
type ListClusterMasterSnapshotTasksResponse struct {

	// **参数解释：** API版本，默认为v3.1 **约束限制：** 不涉及 **取值范围：** - v3.1  **默认取值：** v3.1
	ApiVersion *string `json:"apiVersion,omitempty"`

	// **参数解释：** 任务类型 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Kind *string `json:"kind,omitempty"`

	Metadata *SnapshotTaskMetadata `json:"metadata,omitempty"`

	// **参数解释：** 备份任务列表 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Items *[]SnapshotTask `json:"items,omitempty"`

	Status         *SnapshotTaskStatus `json:"status,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListClusterMasterSnapshotTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterMasterSnapshotTasksResponse struct{}"
	}

	return strings.Join([]string{"ListClusterMasterSnapshotTasksResponse", string(data)}, " ")
}
