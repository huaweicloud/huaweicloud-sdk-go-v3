package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkFlowSpec **参数解释：** 集群升级流程配置信息。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type WorkFlowSpec struct {

	// **参数解释：** 集群ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ClusterID *string `json:"clusterID,omitempty"`

	// **参数解释：** 当前集群版本 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ClusterVersion *string `json:"clusterVersion,omitempty"`

	// **参数解释：** 本次集群升级的目标版本 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	TargetVersion string `json:"targetVersion"`
}

func (o WorkFlowSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkFlowSpec struct{}"
	}

	return strings.Join([]string{"WorkFlowSpec", string(data)}, " ")
}
