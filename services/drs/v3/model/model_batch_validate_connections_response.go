package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchValidateConnectionsResponse Response Object
type BatchValidateConnectionsResponse struct {

	// **参数解释**：  批量测试连接响应体集合，包含每个任务测试连接的结果信息。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Results *[]CheckJobResp `json:"results,omitempty"`

	// **参数解释**：  返回的任务信息总数，表示本次查询成功返回的任务数量，与results数组的长度一致。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchValidateConnectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchValidateConnectionsResponse struct{}"
	}

	return strings.Join([]string{"BatchValidateConnectionsResponse", string(data)}, " ")
}
