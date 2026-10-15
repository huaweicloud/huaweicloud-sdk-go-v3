package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchListJobStatusResponse Response Object
type BatchListJobStatusResponse struct {

	// **参数解释**：  批量查询返回的各任务状态列表，每个元素对应一个查询任务的状态信息。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Results *[]QueryJobStatusResp `json:"results,omitempty"`

	// **参数解释**：  返回的任务信息总数，表示本次查询成功返回的任务数量，与results数组的长度一致。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchListJobStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListJobStatusResponse struct{}"
	}

	return strings.Join([]string{"BatchListJobStatusResponse", string(data)}, " ")
}
