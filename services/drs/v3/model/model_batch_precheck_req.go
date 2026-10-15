package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchPrecheckReq 批量预检查请求体
type BatchPrecheckReq struct {

	// **参数解释**：  批量预检查请求列表，详情请参见PreCheckInfo对象定义。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Jobs []PreCheckInfo `json:"jobs"`
}

func (o BatchPrecheckReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchPrecheckReq struct{}"
	}

	return strings.Join([]string{"BatchPrecheckReq", string(data)}, " ")
}
