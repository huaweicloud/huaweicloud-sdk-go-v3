package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateJobReq 批量创建实时迁移任务请求体
type BatchCreateJobReq struct {

	// **参数解释**：  创建任务请求体列表，详情请参见CreateJobReq对象定义。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Jobs []CreateJobReq `json:"jobs"`
}

func (o BatchCreateJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateJobReq struct{}"
	}

	return strings.Join([]string{"BatchCreateJobReq", string(data)}, " ")
}
