package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ClusterQuotaResource struct {

	// **参数解释**： 资源类型 **约束限制**： 不涉及 **取值范围**： - cluster：Standard/Turbo集群 - autopilot_cluster：Autopilot集群  **默认取值**： 不涉及
	Type *string `json:"type,omitempty"`

	// **参数解释**： 总配额 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Quota *int32 `json:"quota,omitempty"`

	// **参数解释**： 配额已使用数量 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Used *int32 `json:"used,omitempty"`

	// **参数解释**： 单位 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Unit *string `json:"unit,omitempty"`

	// **参数解释**： 配额最小值 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Min *int32 `json:"min,omitempty"`

	// **参数解释**： 配额最大值 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Max *int32 `json:"max,omitempty"`
}

func (o ClusterQuotaResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterQuotaResource struct{}"
	}

	return strings.Join([]string{"ClusterQuotaResource", string(data)}, " ")
}
