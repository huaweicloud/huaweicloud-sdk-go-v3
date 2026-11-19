package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKeyUsageFlowResponse Response Object
type ShowKeyUsageFlowResponse struct {

	// **参数解释：** 使用流图节点列表，当前固定包含一个元素，元素内按类型分组展示各类节点 **取值范围：** 不涉及
	Nodes *[]ShowKeyUsageFlowResponseBodyNodes `json:"nodes,omitempty"`

	// **参数解释：** 使用流关系边列表，表示起点节点到策略节点、策略节点到密钥节点的使用流关系 **取值范围：** 不涉及
	Edges          *[]ShowKeyUsageFlowResponseBodyEdges `json:"edges,omitempty"`
	HttpStatusCode int                                  `json:"-"`
}

func (o ShowKeyUsageFlowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKeyUsageFlowResponse struct{}"
	}

	return strings.Join([]string{"ShowKeyUsageFlowResponse", string(data)}, " ")
}
