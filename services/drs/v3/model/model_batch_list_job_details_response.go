package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchListJobDetailsResponse Response Object
type BatchListJobDetailsResponse struct {

	// **参数解释**：  返回的任务信息总数，表示本次查询成功返回的任务数量，与results数组的长度一致。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Count *int32 `json:"count,omitempty"`

	// **参数解释**：  任务详细信息列表，每个元素为QueryJobResp结构。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Results        *[]QueryJobResp `json:"results,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o BatchListJobDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListJobDetailsResponse struct{}"
	}

	return strings.Join([]string{"BatchListJobDetailsResponse", string(data)}, " ")
}
