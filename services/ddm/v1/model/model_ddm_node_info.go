package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DdmNodeInfo struct {

	// **参数解释**：  节点ID。  **参数范围**：  不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：  状态。  **参数范围**：  不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**：  名称。  **参数范围**：  不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：  可用区编码。  **参数范围**：  不涉及。
	AzCode *string `json:"az_code,omitempty"`

	// **参数解释**：  锁状态。  **参数范围**：  不涉及。
	Actions *[]ActionInfo `json:"actions,omitempty"`

	// **参数解释**：  子网ID。  **参数范围**：  不涉及。
	SubnetId *string `json:"subnet_id,omitempty"`
}

func (o DdmNodeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DdmNodeInfo struct{}"
	}

	return strings.Join([]string{"DdmNodeInfo", string(data)}, " ")
}
