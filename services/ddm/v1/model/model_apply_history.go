package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApplyHistory struct {

	// **参数解释**：  参数组ID。  **参数范围**：  不涉及。
	TargetId *string `json:"target_id,omitempty"`

	// **参数解释**：  参数组的名称。  **参数范围**：  不涉及。
	TargetName *string `json:"target_name,omitempty"`

	// **参数解释**：  应用结果。  **参数范围**：  不涉及。
	ApplyResult *string `json:"apply_result,omitempty"`

	// **参数解释**：  应用日期。  **参数范围**：  不涉及。
	AppliedAt *sdktime.SdkTime `json:"applied_at,omitempty"`

	// **参数解释**：  错误码。  **参数范围**：  不涉及。
	ErrorCode *string `json:"error_code,omitempty"`
}

func (o ApplyHistory) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyHistory struct{}"
	}

	return strings.Join([]string{"ApplyHistory", string(data)}, " ")
}
