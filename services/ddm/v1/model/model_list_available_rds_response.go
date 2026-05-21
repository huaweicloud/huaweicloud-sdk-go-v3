package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAvailableRdsResponse Response Object
type ListAvailableRdsResponse struct {

	// 可用后端DN信息。
	DataNodes *[]AvailableDnInstance `json:"data_nodes,omitempty"`

	// **参数解释**：  分页参数: 起始值。  **参数范围**：   大于等于0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：  分页参数: 每页记录数。  **参数范围**：  大于0且小于等于128。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：  总记录数。  **参数范围**：  不涉及。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAvailableRdsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailableRdsResponse struct{}"
	}

	return strings.Join([]string{"ListAvailableRdsResponse", string(data)}, " ")
}
