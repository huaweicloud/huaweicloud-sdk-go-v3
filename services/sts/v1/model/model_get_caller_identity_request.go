package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetCallerIdentityRequest Request Object
type GetCallerIdentityRequest struct {

	// **参数解释**： 临时安全凭证的security_token字段。  **约束限制**： 通过临时安全凭证调用接口时，需要提供“X-Security-Token”Http头。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`
}

func (o GetCallerIdentityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetCallerIdentityRequest struct{}"
	}

	return strings.Join([]string{"GetCallerIdentityRequest", string(data)}, " ")
}
