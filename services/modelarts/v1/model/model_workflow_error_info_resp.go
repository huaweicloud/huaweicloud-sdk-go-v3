package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkflowErrorInfoResp error info struct
type WorkflowErrorInfoResp struct {

	// **参数解释**：错误码。 **取值范围**：不涉及。
	ErrorCode *string `json:"error_code,omitempty"`

	// **参数解释**：错误信息。 **取值范围**：不涉及。
	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o WorkflowErrorInfoResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowErrorInfoResp struct{}"
	}

	return strings.Join([]string{"WorkflowErrorInfoResp", string(data)}, " ")
}
