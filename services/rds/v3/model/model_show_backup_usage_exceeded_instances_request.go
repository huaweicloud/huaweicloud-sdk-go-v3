package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBackupUsageExceededInstancesRequest Request Object
type ShowBackupUsageExceededInstancesRequest struct {

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn  **默认取值**：  en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// **参数解释**：  查询记录数。  **约束限制**：  不涉及。  **取值范围**：  1-100。  **默认取值**：  100。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：  索引位置，偏移量。  **约束限制**：  不涉及。  **取值范围**：  0及以上。  **默认取值**：  0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ShowBackupUsageExceededInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackupUsageExceededInstancesRequest struct{}"
	}

	return strings.Join([]string{"ShowBackupUsageExceededInstancesRequest", string(data)}, " ")
}
