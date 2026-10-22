package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnitGroupPolicy **参数解释：** 单副本资源实例调度约束配置。 **约束限制：** 不涉及。 **取值范围：** 不涉及。
type UnitGroupPolicy struct {

	// **参数解释：** 调度范围，仅A5超节点规格生效。 **约束限制：** - 仅当资源池包含A5(snt9e)规格时有效，非A5规格不执行调度范围校验。 - A5规格仅支持LWS工作负载类型，不支持DEPLOYMENT。 - A5规格需要Volcano调度插件版本满足要求。 **取值范围：** - HYPERNODE：超节点调度 - SLICE：柜级调度 **默认取值：** 不涉及。
	ScheduleScope *string `json:"schedule_scope,omitempty"`
}

func (o UnitGroupPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnitGroupPolicy struct{}"
	}

	return strings.Join([]string{"UnitGroupPolicy", string(data)}, " ")
}
