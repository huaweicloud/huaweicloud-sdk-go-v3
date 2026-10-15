package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchQueryJobReqPage 带分页的批量查询任务详情请求体
type BatchQueryJobReqPage struct {

	// **参数解释**：  批量查询任务ID列表，通过该参数指定需要查询的一个或多个任务。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Jobs []string `json:"jobs"`

	PageReq *PageReq `json:"page_req,omitempty"`
}

func (o BatchQueryJobReqPage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchQueryJobReqPage struct{}"
	}

	return strings.Join([]string{"BatchQueryJobReqPage", string(data)}, " ")
}
