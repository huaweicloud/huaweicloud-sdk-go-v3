package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteReadonlyNodesRequestBody struct {

	// **参数解释**: 即将删除的只读节点ID列表。 **约束限制**: 不涉及。 **取值范围**： 不涉及。 **默认取值**: 不涉及。
	NodeIds []string `json:"node_ids"`
}

func (o DeleteReadonlyNodesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteReadonlyNodesRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteReadonlyNodesRequestBody", string(data)}, " ")
}
