package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchColdTableStorageResponse Response Object
type SwitchColdTableStorageResponse struct {

	// **参数解释**：  操作结果。  **取值范围**：  - SUCCESS：成功。 - FAILED：失败。
	Result *string `json:"result,omitempty"`

	// **参数解释**：  校验失败的表信息。  当result取值为FAILED时，返回此参数。
	FailedItems    *[]FailedItem `json:"failed_items,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o SwitchColdTableStorageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchColdTableStorageResponse struct{}"
	}

	return strings.Join([]string{"SwitchColdTableStorageResponse", string(data)}, " ")
}
