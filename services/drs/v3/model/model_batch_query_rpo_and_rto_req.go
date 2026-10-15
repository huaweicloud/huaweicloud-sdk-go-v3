package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchQueryRpoAndRtoReq 批量查询RPO和RTO的请求体
type BatchQueryRpoAndRtoReq struct {

	// **参数解释**：  批量查询RPO和RTO的任务ID列表，通过该参数指定需要查询的一个或多个任务。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Jobs []string `json:"jobs"`
}

func (o BatchQueryRpoAndRtoReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchQueryRpoAndRtoReq struct{}"
	}

	return strings.Join([]string{"BatchQueryRpoAndRtoReq", string(data)}, " ")
}
