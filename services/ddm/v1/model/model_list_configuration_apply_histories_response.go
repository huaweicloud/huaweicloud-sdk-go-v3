package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConfigurationApplyHistoriesResponse Response Object
type ListConfigurationApplyHistoriesResponse struct {

	// **参数解释**：  参数组应用记录相关信息的集合。  **参数范围**：  不涉及。
	Histories *[]ApplyHistory `json:"histories,omitempty"`

	// **参数解释**：  分页参数: 每页记录数。  **参数范围**：  不涉及。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListConfigurationApplyHistoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigurationApplyHistoriesResponse struct{}"
	}

	return strings.Join([]string{"ListConfigurationApplyHistoriesResponse", string(data)}, " ")
}
