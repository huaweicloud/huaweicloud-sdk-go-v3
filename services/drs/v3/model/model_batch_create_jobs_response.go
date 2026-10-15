package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateJobsResponse Response Object
type BatchCreateJobsResponse struct {

	// **参数解释**：  批量创建任务的响应体集合，详情请参见CreateJobResp对象定义。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Results *[]CreateJobResp `json:"results,omitempty"`

	// **参数解释**：  返回的任务信息总数，表示本次查询成功返回的任务数量，与results数组的长度一致。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchCreateJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateJobsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateJobsResponse", string(data)}, " ")
}
