package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetCallerIdentityResponse Response Object
type GetCallerIdentityResponse struct {

	// **参数解释**： 账号ID。  **取值范围**： 不涉及。
	AccountId *string `json:"account_id,omitempty"`

	// **参数解释**： 主体URN。  **取值范围**： 不涉及。
	PrincipalUrn *string `json:"principal_urn,omitempty"`

	// **参数解释**： 主体ID。  **取值范围**： 不涉及。
	PrincipalId    *string `json:"principal_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o GetCallerIdentityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetCallerIdentityResponse struct{}"
	}

	return strings.Join([]string{"GetCallerIdentityResponse", string(data)}, " ")
}
