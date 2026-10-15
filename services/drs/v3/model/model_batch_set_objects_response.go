package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSetObjectsResponse Response Object
type BatchSetObjectsResponse struct {

	// **参数解释**：  批量对象选择操作的任务总数，表示本次请求中处理的对象选择任务总数量。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	AllCounts *int64 `json:"all_counts,omitempty"`

	// **参数解释**：  批量对象选择操作的响应列表，详情请参见DatabaseObjectResp对象定义。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Results        *[]DatabaseObjectResp `json:"results,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o BatchSetObjectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetObjectsResponse struct{}"
	}

	return strings.Join([]string{"BatchSetObjectsResponse", string(data)}, " ")
}
