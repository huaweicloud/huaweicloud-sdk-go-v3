package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDiagnosisStackResponse Response Object
type ShowDiagnosisStackResponse struct {

	// **参数解释**： 线程名。 **取值范围**： 不涉及。
	ThreadName *string `json:"thread_name,omitempty"`

	// **参数解释**： 堆信息。 **取值范围**： 不涉及。
	Stack          *string `json:"stack,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDiagnosisStackResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDiagnosisStackResponse struct{}"
	}

	return strings.Join([]string{"ShowDiagnosisStackResponse", string(data)}, " ")
}
