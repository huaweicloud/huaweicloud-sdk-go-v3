package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRebootPoolNodesResponse Response Object
type BatchRebootPoolNodesResponse struct {

	// **参数解释**：节点名称集合。 **约束限制**：不涉及。
	NodeNames      []string `json:"nodeNames"`
	HttpStatusCode int      `json:"-"`
}

func (o BatchRebootPoolNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRebootPoolNodesResponse struct{}"
	}

	return strings.Join([]string{"BatchRebootPoolNodesResponse", string(data)}, " ")
}
