package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeRisks 节点风险来源
type NodeRisks struct {

	// **参数解释：** 用户节点ID。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	NodeID *string `json:"NodeID,omitempty"`
}

func (o NodeRisks) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeRisks struct{}"
	}

	return strings.Join([]string{"NodeRisks", string(data)}, " ")
}
