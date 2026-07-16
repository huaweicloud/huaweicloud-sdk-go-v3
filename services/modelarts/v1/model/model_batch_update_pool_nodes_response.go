package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdatePoolNodesResponse Response Object
type BatchUpdatePoolNodesResponse struct {

	// **参数解释**：更新成功的节点名称列表。
	SuccessNodeNames *[]string `json:"successNodeNames,omitempty"`

	// **参数解释**：更新失败的节点名称列表。
	FailNodeNames  *[]string `json:"failNodeNames,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchUpdatePoolNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdatePoolNodesResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdatePoolNodesResponse", string(data)}, " ")
}
