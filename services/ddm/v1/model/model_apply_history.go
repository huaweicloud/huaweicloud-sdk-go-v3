package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApplyHistory struct {

	// **参数解释**：  参数组ID。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，长度为36个字符。  **默认取值**：  不涉及。
	TargetId *string `json:"target_id,omitempty"`

	// **参数解释**：  参数组的名称。  **参数范围**：  不涉及。
	TargetName *string `json:"target_name,omitempty"`

	// **参数解释**：  应用结果。  **参数范围**：  不涉及。
	ApplyResult *string `json:"apply_result,omitempty"`

	// **参数解释**：  应用日期。  格式为yyyy-mm-dd Thh:mm:ssZ。  其中，T指定某个时间的开始；Z指示 UTC 时间。  **参数范围**：  不涉及。
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
