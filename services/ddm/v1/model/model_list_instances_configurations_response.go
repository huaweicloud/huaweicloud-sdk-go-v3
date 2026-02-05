package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstancesConfigurationsResponse Response Object
type ListInstancesConfigurationsResponse struct {

	// **参数解释**：  查询可应用的实例列表返回相关信息的集合。  **参数范围**：  不涉及。
	Entities *[]ApplicableInstance `json:"entities,omitempty"`

	// **参数解释**：  分页参数: 每页记录数。  **参数范围**：  不涉及。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstancesConfigurationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesConfigurationsResponse struct{}"
	}

	return strings.Join([]string{"ListInstancesConfigurationsResponse", string(data)}, " ")
}
