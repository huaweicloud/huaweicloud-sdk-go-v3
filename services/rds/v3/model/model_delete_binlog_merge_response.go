package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBinlogMergeResponse Response Object
type DeleteBinlogMergeResponse struct {

	// **参数解释**：  删除操作执行状态。  **约束限制**：  不涉及。  **取值范围**：  - COMPLETED (已完成)  **默认取值**：  不涉及。
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteBinlogMergeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBinlogMergeResponse struct{}"
	}

	return strings.Join([]string{"DeleteBinlogMergeResponse", string(data)}, " ")
}
