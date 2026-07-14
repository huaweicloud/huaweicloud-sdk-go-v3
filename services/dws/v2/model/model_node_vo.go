package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeVo **参数解释**： 节点实例信息。 **取值范围**： 不涉及。
type NodeVo struct {

	// **参数解释**： 节点实例ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 节点实例名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 节点状态。 **取值范围**： - ACTIVE：正常。 - FAILED：不可用。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 节点类型。 **取值范围**： - cn：协调节点。 - dn：数据节点。
	InstType *string `json:"inst_type,omitempty"`
}

func (o NodeVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeVo struct{}"
	}

	return strings.Join([]string{"NodeVo", string(data)}, " ")
}
