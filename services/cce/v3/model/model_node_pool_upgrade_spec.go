package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodePoolUpgradeSpec 同步点池请求详细参数
type NodePoolUpgradeSpec struct {

	// **参数解释**： 每批最大同步节点。节点升级时，允许节点不可用的最大数量。节点重置方式进行同步时节点将不可用，请合理设置该参数，尽量避免出现集群节点不可用数量过多导致Pod无法调度的情况。 **约束限制**： 不涉及 **取值范围**： 取值范围[1-20] **默认取值**： 不涉及
	MaxUnavailable *int32 `json:"maxUnavailable,omitempty"`

	// Pod无法驱逐时，是否强制重置。
	Force *bool `json:"force,omitempty"`

	RetryTimes *int32 `json:"retryTimes,omitempty"`

	SkippedNodes *[]string `json:"skippedNodes,omitempty"`

	// **参数解释**： 本次操作同步的节点池中选择的节点ID列表，获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	NodeIDs *[]string `json:"nodeIDs,omitempty"`

	// **参数解释**： 节点池ID，获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	NodePoolID string `json:"nodePoolID"`

	NodeTemplate *NodeTemplate `json:"nodeTemplate,omitempty"`

	// **参数解释**： 是否在同步节点池任务下发前校验节点storage参数。如果开启校验，当存在节点storage参数异常时接口将直接返回错误；如果不开启校验，接口将先下发同步节点池任务，当存在节点storage参数异常时，同步节点池任务将失败。 **约束限制**： 当开启节点storage参数校验时，每次同步的节点池节点个数不能大于200个。 **取值范围**： - false：不校验节点storage参数 - true：校验节点storage参数  **默认取值**： false
	ValidateStorage *bool `json:"validateStorage,omitempty"`
}

func (o NodePoolUpgradeSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodePoolUpgradeSpec struct{}"
	}

	return strings.Join([]string{"NodePoolUpgradeSpec", string(data)}, " ")
}
