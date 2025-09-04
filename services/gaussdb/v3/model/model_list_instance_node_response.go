package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceNodeResponse Response Object
type ListInstanceNodeResponse struct {

	// **参数解释**： HTAP标准版实例节点列表。  **约束限制**： 不涉及。
	NodeList       *[]HtapNodeInfoResponseBodyNodeList `json:"node_list,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o ListInstanceNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceNodeResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceNodeResponse", string(data)}, " ")
}
