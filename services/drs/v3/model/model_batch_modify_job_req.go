package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchModifyJobReq 批量修改任务请求体。
type BatchModifyJobReq struct {

	// **参数解释**：  修改任务请求体列表，包含需要修改的一个或多个任务信息，详情请参见ModifyJobReq对象定义。  **约束限制**：  - 建议批量调用数量不超过10个。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Jobs []ModifyJobReq `json:"jobs"`
}

func (o BatchModifyJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchModifyJobReq struct{}"
	}

	return strings.Join([]string{"BatchModifyJobReq", string(data)}, " ")
}
