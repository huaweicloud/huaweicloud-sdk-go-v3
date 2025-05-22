package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteClusterNodesRequestBody **参数解释**： 节点操作请求。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type DeleteClusterNodesRequestBody struct {

	// **参数解释**： 空闲节点ID列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	NodeList []string `json:"node_list"`

	// **参数解释**： 操作类型，一般传delete即可。 **约束限制**： 不涉及。 **取值范围**： clear：清理创建失败的空闲节点 delete：删除空闲节点 **默认取值**： 不涉及。
	OperateType string `json:"operate_type"`
}

func (o DeleteClusterNodesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteClusterNodesRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteClusterNodesRequestBody", string(data)}, " ")
}
