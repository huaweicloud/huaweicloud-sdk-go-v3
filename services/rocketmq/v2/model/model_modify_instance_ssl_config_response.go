package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyInstanceSslConfigResponse Response Object
type ModifyInstanceSslConfigResponse struct {

	// **参数解释**： 任务ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释**： 协议模式。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	TlsMode        *string `json:"tls_mode,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ModifyInstanceSslConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyInstanceSslConfigResponse struct{}"
	}

	return strings.Join([]string{"ModifyInstanceSslConfigResponse", string(data)}, " ")
}
