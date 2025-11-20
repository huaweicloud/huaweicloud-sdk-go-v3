package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeClusterRequestBodyExtendParam **参数解释**： 变更集群规格扩展字段 **约束限制**： 不涉及
type ResizeClusterRequestBodyExtendParam struct {

	// **参数解释**： 专属云CCE集群可指定控制节点的规格 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	DecMasterFlavor *string `json:"decMasterFlavor,omitempty"`

	// **参数解释**： 是否自动扣款 **约束限制**： 不涉及 **取值范围**： - “true”：自动扣款 - “false”：不自动扣款 > 包周期集群时生效，不填写此参数时默认不会自动扣款。  **默认取值**： 不涉及
	IsAutoPay *string `json:"isAutoPay,omitempty"`
}

func (o ResizeClusterRequestBodyExtendParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeClusterRequestBodyExtendParam struct{}"
	}

	return strings.Join([]string{"ResizeClusterRequestBodyExtendParam", string(data)}, " ")
}
