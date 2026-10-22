package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnitGroupPolicyResponse **参数解释：** 单副本资源实例调度约束配置。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
type UnitGroupPolicyResponse struct {

	// **参数解释：** 调度范围。 **取值范围：** HYPERNODE：超节点调度 SLICE：柜级调度
	ScheduleScope *string `json:"schedule_scope,omitempty"`
}

func (o UnitGroupPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnitGroupPolicyResponse struct{}"
	}

	return strings.Join([]string{"UnitGroupPolicyResponse", string(data)}, " ")
}
