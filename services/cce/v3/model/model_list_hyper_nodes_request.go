package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHyperNodesRequest Request Object
type ListHyperNodesRequest struct {

	// 集群ID，获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。
	ClusterId string `json:"cluster_id"`

	// **参数解释**： 设置每页显示的数据条数 **约束限制**： 不涉及 **取值范围**： 1到1000之间（含1和1000）的整数 **默认取值**： 100
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 设置从第几条数据开始显示（用于翻页），比如输入0表示从第一条数据开始，输入10表示跳过前10条，从第11条开始显示，不填时默认从第一条开始显示（即默认为0）。 **约束限制**： 不涉及 **取值范围**： 32 位非负整数 **默认取值**： 0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListHyperNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHyperNodesRequest struct{}"
	}

	return strings.Join([]string{"ListHyperNodesRequest", string(data)}, " ")
}
