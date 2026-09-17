package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterCheckStatus **参数解释：** 集群限制检查状态 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type ClusterCheckStatus struct {

	// **参数解释：** 状态 **约束限制：** 不涉及 **取值范围：** - Init：初始化 - Running：运行中 - Success：成功 - Failed：失败  **默认取值：** 不涉及
	Phase *string `json:"phase,omitempty"`

	// **参数解释：** 检查项状态集合 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ItemsStatus *[]PreCheckItemStatus `json:"itemsStatus,omitempty"`
}

func (o ClusterCheckStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterCheckStatus struct{}"
	}

	return strings.Join([]string{"ClusterCheckStatus", string(data)}, " ")
}
