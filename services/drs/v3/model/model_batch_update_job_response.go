package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateJobResponse Response Object
type BatchUpdateJobResponse struct {

	// **参数解释**：  返回的任务信息总数，表示本次修改成功返回的任务数量，与results数组的长度一致。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Count *int32 `json:"count,omitempty"`

	// **参数解释**：  批量修改任务返回列表，包含每个任务的修改结果，详情请参见ModifyJobResp对象定义。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Results        *[]ModifyJobResp `json:"results,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o BatchUpdateJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateJobResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateJobResponse", string(data)}, " ")
}
