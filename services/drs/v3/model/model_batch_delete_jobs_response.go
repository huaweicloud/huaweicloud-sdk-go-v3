package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteJobsResponse Response Object
type BatchDeleteJobsResponse struct {

	// **参数解释**：  批量结束任务或删除任务的响应体集合，包含每个任务的操作结果。详情请参见DeleteJobResp对象定义。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Results *[]DeleteJobResp `json:"results,omitempty"`

	// **参数解释**：  批量结束或删除任务的结果总数，与results数组的长度一致。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchDeleteJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteJobsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteJobsResponse", string(data)}, " ")
}
