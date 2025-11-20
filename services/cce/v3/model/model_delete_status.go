package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteStatus
type DeleteStatus struct {

	// **参数解释**： 集群删除时已经存在的集群资源记录总数 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	PreviousTotal *int32 `json:"previous_total,omitempty"`

	// **参数解释**： 基于当前集群资源记录信息，生成实际最新资源记录总数 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	CurrentTotal *int32 `json:"current_total,omitempty"`

	// **参数解释**： 集群删除时更新的资源记录总数 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Updated *int32 `json:"updated,omitempty"`

	// **参数解释**： 集群删除时更新的资源记录总数 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Added *int32 `json:"added,omitempty"`

	// **参数解释**： 集群删除时删除的资源记录总数 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Deleted *int32 `json:"deleted,omitempty"`
}

func (o DeleteStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStatus struct{}"
	}

	return strings.Join([]string{"DeleteStatus", string(data)}, " ")
}
