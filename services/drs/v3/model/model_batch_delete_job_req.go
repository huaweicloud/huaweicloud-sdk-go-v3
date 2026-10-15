package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteJobReq **参数解释**：  批量结束与删除任务请求体，包含待操作的任务列表。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
type BatchDeleteJobReq struct {

	// **参数解释**：  批量结束与删除任务请求列表，指定需要结束或删除的任务信息。详情请参见DeleteJobReq对象定义。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Jobs []DeleteJobReq `json:"jobs"`
}

func (o BatchDeleteJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteJobReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteJobReq", string(data)}, " ")
}
