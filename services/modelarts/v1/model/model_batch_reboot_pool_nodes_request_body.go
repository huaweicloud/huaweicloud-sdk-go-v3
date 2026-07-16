package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRebootPoolNodesRequestBody 重启节点名称集合 { \"nodeNames\": [   \"os-node-created-vrvrq\",   \"os-node-created-4jczv\"  ] }
type BatchRebootPoolNodesRequestBody struct {

	// **参数解释**：节点名称集合。 **约束限制**：不涉及。
	NodeNames []string `json:"nodeNames"`
}

func (o BatchRebootPoolNodesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRebootPoolNodesRequestBody struct{}"
	}

	return strings.Join([]string{"BatchRebootPoolNodesRequestBody", string(data)}, " ")
}
