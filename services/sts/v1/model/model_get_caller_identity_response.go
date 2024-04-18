package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetCallerIdentityResponse Response Object
type GetCallerIdentityResponse struct {

	// 华为云账号ID。
	AccountId *string `json:"account_id,omitempty"`

	// 主体URN。
	PrincipalUrn *string `json:"principal_urn,omitempty"`

	// 主体ID。
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
