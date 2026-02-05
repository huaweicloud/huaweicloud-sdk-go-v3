package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchModifyReadWriteStrategyResponse Response Object
type BatchModifyReadWriteStrategyResponse struct {

	// **参数解释**：  是否成功。  **参数范围**：  不涉及。
	Success *bool `json:"success,omitempty"`

	// **参数解释**：  请求id。  **参数范围**：  不涉及。
	RequestId *string `json:"request_id,omitempty"`

	// **参数解释**：  返回信息。  **参数范围**：  不涉及。
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o BatchModifyReadWriteStrategyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchModifyReadWriteStrategyResponse struct{}"
	}

	return strings.Join([]string{"BatchModifyReadWriteStrategyResponse", string(data)}, " ")
}
