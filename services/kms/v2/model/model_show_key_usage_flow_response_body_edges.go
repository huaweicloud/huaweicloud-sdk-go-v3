package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowKeyUsageFlowResponseBodyEdges struct {

	// **参数解释：** 边起点节点ID **取值范围：** 不涉及
	Source *string `json:"source,omitempty"`

	// **参数解释：** 边终点节点ID **取值范围：** 不涉及
	Target *string `json:"target,omitempty"`
}

func (o ShowKeyUsageFlowResponseBodyEdges) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKeyUsageFlowResponseBodyEdges struct{}"
	}

	return strings.Join([]string{"ShowKeyUsageFlowResponseBodyEdges", string(data)}, " ")
}
