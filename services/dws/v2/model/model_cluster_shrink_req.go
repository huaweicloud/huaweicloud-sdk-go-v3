package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterShrinkReq **参数解释**： 逻辑集群缩容请求体。 **约束限制**： 该值不能为空。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type ClusterShrinkReq struct {

	// **参数解释**：  缩容节点个数。  **约束限制**：  该值不能为空。  **取值范围**：  大于0的整数。  **默认取值**：  不涉及。
	ShrinkNumber *int32 `json:"shrink_number,omitempty"`

	// **参数解释**：  是否是在线缩容。  **约束限制**：  不涉及。  **取值范围**：  false|true。  **默认取值**：  false。
	Online *bool `json:"online,omitempty"`

	// **参数解释**：  是否是缩容失败后重试。  **约束限制**：  不涉及。  **取值范围**：  false|true。  **默认取值**：  false。
	Retry *bool `json:"retry,omitempty"`

	// **参数解释**：  是否需要委托。  **约束限制**：  不涉及。  **取值范围**：  false或true。  **默认取值**：  false。
	NeedAgency *bool `json:"need_agency,omitempty"`

	// **参数解释**：  重分布并发配置数。  **约束限制**：  不涉及。  **取值范围**：  1~200。  **默认取值**：  4。
	ParallelJobs *int32 `json:"parallel_jobs,omitempty"`

	// **参数解释**：  类型字段，字段已废弃不再生效。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**：  操作前是否执行备份，字段已废弃不再生效。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	ForceBackup *bool `json:"force_backup,omitempty"`
}

func (o ClusterShrinkReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterShrinkReq struct{}"
	}

	return strings.Join([]string{"ClusterShrinkReq", string(data)}, " ")
}
