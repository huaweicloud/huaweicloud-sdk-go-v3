package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchListProgressesResponse Response Object
type BatchListProgressesResponse struct {

	// **参数解释**：  返回的任务进度信息总数，表示本次查询成功返回的任务数量，与results数组的长度一致。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Count *int32 `json:"count,omitempty"`

	// **参数解释**：  批量查询迁移进度的响应体集合，详情请参见QueryProgressResp定义。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Results        *[]QueryProgressResp `json:"results,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o BatchListProgressesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListProgressesResponse struct{}"
	}

	return strings.Join([]string{"BatchListProgressesResponse", string(data)}, " ")
}
