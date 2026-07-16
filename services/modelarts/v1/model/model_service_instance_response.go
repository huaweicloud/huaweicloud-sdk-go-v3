package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServiceInstanceResponse 服务实例列表。
type ServiceInstanceResponse struct {

	// **参数解释：** 服务实例名字。 **取值范围：** 不涉及。
	InstanceName *string `json:"instance_name,omitempty"`

	// **参数解释：** 服务实例状态。 **取值范围：** - RUNNING：运行中 - PENDING：未就绪 - CONCERNING：告警 - FAILED：失败 - UNKNOWN：未知 - DELETED：已删除
	Status *string `json:"status,omitempty"`

	// **参数解释：** 服务实例权重。 **取值范围：** [0, 100] 或者为空。
	Weight *int32 `json:"weight,omitempty"`

	// **参数解释：** 服务实例pod数量。 **取值范围：** 不涉及。
	PodCount *int64 `json:"pod_count,omitempty"`

	// **参数解释：** 服务实例运行中pod数量。 **取值范围：** 不涉及。
	RunningPodCount *int64 `json:"running_pod_count,omitempty"`

	// **参数解释：** 服务实例最近更新时间。 **取值范围：** 不涉及。
	UpdateAt *int64 `json:"update_at,omitempty"`
}

func (o ServiceInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceInstanceResponse struct{}"
	}

	return strings.Join([]string{"ServiceInstanceResponse", string(data)}, " ")
}
