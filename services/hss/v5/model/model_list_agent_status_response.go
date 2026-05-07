package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAgentStatusResponse Response Object
type ListAgentStatusResponse struct {

	// **参数解释**: 记录总数 **取值范围**: 最小值0，最大值10
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: Agent状态列表 **取值范围**: 最小值0，最大值10
	DataList       *[]AgentStatusInfo `json:"data_list,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListAgentStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAgentStatusResponse struct{}"
	}

	return strings.Join([]string{"ListAgentStatusResponse", string(data)}, " ")
}
