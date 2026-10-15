package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCheckResultsResponse Response Object
type BatchCheckResultsResponse struct {

	// **参数解释**：  批量查询预检查结果响应体集合，包含每个任务的预检查结果详情，详情请参见QueryPreCheckResp对象定义。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Results *[]QueryPreCheckResp `json:"results,omitempty"`

	// **参数解释**：  返回的任务信息总数，表示本次查询成功返回的任务数量，与results数组的长度一致。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchCheckResultsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCheckResultsResponse struct{}"
	}

	return strings.Join([]string{"BatchCheckResultsResponse", string(data)}, " ")
}
