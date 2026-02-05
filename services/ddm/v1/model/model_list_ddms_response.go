package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDdmsResponse Response Object
type ListDdmsResponse struct {

	// **参数解释**：  实例数量。  **参数范围**：  不涉及。
	InstanceNum *int32 `json:"instance_num,omitempty"`

	// **参数解释**：  实例列表的集合。  **参数范围**：  不涉及。
	Instances *[]CustomerInstanceVo `json:"instances,omitempty"`

	// **参数解释**：  分页大小。  **参数范围**：  不涉及。
	PageSize *int32 `json:"page_size,omitempty"`

	// **参数解释**：  实例总量。  **参数范围**：  不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**：  总分页数。  **参数范围**：  不涉及。
	TotalPage *int32 `json:"total_page,omitempty"`

	// **参数解释**：  分页序号。  **参数范围**：  不涉及。
	PageNo         *int32 `json:"page_no,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDdmsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDdmsResponse struct{}"
	}

	return strings.Join([]string{"ListDdmsResponse", string(data)}, " ")
}
