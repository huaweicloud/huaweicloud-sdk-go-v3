package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateBackupPolicyResponse Response Object
type BatchUpdateBackupPolicyResponse struct {

	// **参数解释**：  备份策略设置异常信息，当所有实例设置成功，该值是空。
	FailedResults *[]UpdateBackupPolicyResult `json:"failed_results,omitempty"`

	// **参数解释**：  设置成功的实例数量。  **取值范围**：  0-50。
	SuccessCount *int32 `json:"success_count,omitempty"`

	// **参数解释**：  设置失败的实例数量。  **取值范围**：  0-50。
	FailedCount    *int32 `json:"failed_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchUpdateBackupPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateBackupPolicyResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateBackupPolicyResponse", string(data)}, " ")
}
