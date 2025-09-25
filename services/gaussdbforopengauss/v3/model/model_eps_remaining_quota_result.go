package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EpsRemainingQuotaResult struct {

	// **参数解释**: 企业项目ID。 **取值范围**: 不涉及。
	EpsTag *string `json:"eps_tag,omitempty"`

	// **参数解释**: 实例配额。 **取值范围**: 不涉及。
	InstanceEpsQuota *int32 `json:"instance_eps_quota,omitempty"`

	// **参数解释**: CPU配额。 **取值范围**: 不涉及。
	CpuEpsQuota *int32 `json:"cpu_eps_quota,omitempty"`

	// **参数解释**: 内存配额。 **取值范围**: 不涉及。
	MemEpsQuota *int32 `json:"mem_eps_quota,omitempty"`

	// **参数解释**: 存储空间配额。 **取值范围**: 不涉及。
	VolumeEpsQuota *int32 `json:"volume_eps_quota,omitempty"`

	// **参数解释**: 实例剩余配额。 **取值范围**: 不涉及。
	InstanceEpsRemainingQuota *int32 `json:"instance_eps_remaining_quota,omitempty"`

	// **参数解释**: CPU剩余配额。 **取值范围**: 不涉及。
	CpuEpsRemainingQuota *int32 `json:"cpu_eps_remaining_quota,omitempty"`

	// **参数解释**: 内存剩余配额。 **取值范围**: 不涉及。
	MemEpsRemainingQuota *int32 `json:"mem_eps_remaining_quota,omitempty"`

	// **参数解释**: 存储空间剩余配额。 **取值范围**: 不涉及。
	VolumeEpsRemainingQuota *int32 `json:"volume_eps_remaining_quota,omitempty"`
}

func (o EpsRemainingQuotaResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EpsRemainingQuotaResult struct{}"
	}

	return strings.Join([]string{"EpsRemainingQuotaResult", string(data)}, " ")
}
