package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogicalClusterInfo **参数解释**： 逻辑集群对象。 **取值范围**： 不涉及。
type LogicalClusterInfo struct {

	// **参数解释**： 逻辑集群ID。 **取值范围**： 不涉及。
	LogicalClusterId *string `json:"logical_cluster_id,omitempty"`

	// **参数解释**： 逻辑集群名称。 **取值范围**： 不涉及。
	LogicalClusterName *string `json:"logical_cluster_name,omitempty"`

	// **参数解释**： 逻辑集群主机信息。 **取值范围**： 不涉及。
	ClusterRings *[]ClusterRing `json:"cluster_rings,omitempty"`

	// **参数解释**： 逻辑集群状态。 **取值范围**： - Failed：失败。 - Normal: 正常。 - Unavailable：不可用。 - Redistributing：重分布中。 - Redistribute_failed：重分布失败。 - Deleted：已删除。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 是否为第一个逻辑集群。历史版本中第1个创建或者转换的逻辑集群不能删除，因为其中包含了一些系统视图。 **取值范围**： 不涉及。
	FirstLogicalCluster *bool `json:"first_logical_cluster,omitempty"`

	ActionInfo *ActionInfo `json:"action_info,omitempty"`

	// **参数解释**： 是否允许编辑。 **取值范围**： 不涉及。
	EditEnable *bool `json:"edit_enable,omitempty"`

	// **参数解释**： 是否允许重启。 **取值范围**： 不涉及。
	RestartEnable *bool `json:"restart_enable,omitempty"`

	// **参数解释**： 是否允许删除。 **取值范围**： 不涉及。
	DeleteEnable *bool `json:"delete_enable,omitempty"`

	// **参数解释**： 是否允许弹性伸缩。 **取值范围**： 不涉及。
	AddToElastic *bool `json:"add_to_elastic,omitempty"`

	// **参数解释**： 逻辑集群模式。 **取值范围**： 不涉及。
	Mode *string `json:"mode,omitempty"`

	// **参数解释**： 作业等待时间。 **取值范围**： 不涉及。
	WaitingForKilling *int32 `json:"waiting_for_killing,omitempty"`

	// **参数解释**： 逻辑集群类型。 **取值范围**： - createFromPlan：计划弹性。 - createFromElastic：自动弹性。
	ClusterType *string `json:"cluster_type,omitempty"`
}

func (o LogicalClusterInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogicalClusterInfo struct{}"
	}

	return strings.Join([]string{"LogicalClusterInfo", string(data)}, " ")
}
