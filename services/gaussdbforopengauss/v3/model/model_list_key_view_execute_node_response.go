package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListKeyViewExecuteNodeResponse Response Object
type ListKeyViewExecuteNodeResponse struct {

	// **参数解释**: 节点信息列表。
	Nodes          *[]OpsKeyViewExecuteNodeResult `json:"nodes,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ListKeyViewExecuteNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKeyViewExecuteNodeResponse struct{}"
	}

	return strings.Join([]string{"ListKeyViewExecuteNodeResponse", string(data)}, " ")
}
