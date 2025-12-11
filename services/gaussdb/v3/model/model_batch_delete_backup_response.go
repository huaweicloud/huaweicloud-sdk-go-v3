package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteBackupResponse Response Object
type BatchDeleteBackupResponse struct {

	// **参数解释**：  备份删除异常信息。  **取值范围**：  当所有备份删除成功，该值是空。
	FailedResults *[]DeleteBackupResult `json:"failed_results,omitempty"`

	// **参数解释**：  删除成功的数量。  **取值范围**：  0-50。
	SuccessCount *int32 `json:"success_count,omitempty"`

	// **参数解释**：  删除失败的数量。  **取值范围**：  0-50。
	FailedCount    *int32 `json:"failed_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchDeleteBackupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteBackupResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteBackupResponse", string(data)}, " ")
}
