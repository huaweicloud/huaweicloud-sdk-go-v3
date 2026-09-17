package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeWorkFlowUpdateRequestBodyStatus **参数解释：** 集群升级流程的执行状态 **约束限制：** 当前仅支持Cancel **取值范围：** - Cancel：表示取消升级  **默认取值：** 不涉及
type UpgradeWorkFlowUpdateRequestBodyStatus struct {
	Phase *WorkFlowPhase `json:"phase,omitempty"`
}

func (o UpgradeWorkFlowUpdateRequestBodyStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeWorkFlowUpdateRequestBodyStatus struct{}"
	}

	return strings.Join([]string{"UpgradeWorkFlowUpdateRequestBodyStatus", string(data)}, " ")
}
