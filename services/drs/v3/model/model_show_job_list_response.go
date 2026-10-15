package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobListResponse Response Object
type ShowJobListResponse struct {

	// **参数解释**：  任务总数，表示满足查询条件的任务记录总数量。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	TotalRecord *int32 `json:"total_record,omitempty"`

	// **参数解释**：  任务信息列表，包含满足查询条件的所有任务详情。详情请参见JobInfo对象定义。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Jobs           *[]JobInfo `json:"jobs,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ShowJobListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobListResponse struct{}"
	}

	return strings.Join([]string{"ShowJobListResponse", string(data)}, " ")
}
