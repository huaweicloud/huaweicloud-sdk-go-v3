package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchStartJobsResponse Response Object
type BatchStartJobsResponse struct {

	// **参数解释**：  批量启动任务返回列表，包含每个任务的启动结果信息。详情请参见StartJobResp对象定义。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Results *[]StartJobResp `json:"results,omitempty"`

	// **参数解释**：  返回的任务信息总数，表示本次批量启动操作返回的任务结果数量，与results数组的长度一致。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchStartJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStartJobsResponse struct{}"
	}

	return strings.Join([]string{"BatchStartJobsResponse", string(data)}, " ")
}
