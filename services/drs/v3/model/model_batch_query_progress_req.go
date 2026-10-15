package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchQueryProgressReq 批量查询同步进度的请求体
type BatchQueryProgressReq struct {

	// **参数解释**：  批量查询任务ID列表，通过该参数指定需要查询进度的一个或多个任务。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Jobs []string `json:"jobs"`
}

func (o BatchQueryProgressReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchQueryProgressReq struct{}"
	}

	return strings.Join([]string{"BatchQueryProgressReq", string(data)}, " ")
}
