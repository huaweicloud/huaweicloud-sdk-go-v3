package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ErrorInfo **参数解释**： 报错详情。 **取值范围**： 不涉及。
type ErrorInfo struct {

	// **参数解释**： 错误码。 **取值范围**： 不涉及。
	Code *string `json:"code,omitempty"`

	// **参数解释**： 错误信息。 **取值范围**： 不涉及。
	Message *string `json:"message,omitempty"`
}

func (o ErrorInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ErrorInfo struct{}"
	}

	return strings.Join([]string{"ErrorInfo", string(data)}, " ")
}
