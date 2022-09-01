package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// scope信息
type GetIdTokenIdScopeBody struct {
	Domain *GetIdTokenScopeDomainOrProjectBody `json:"domain,omitempty" xml:"domain"`

	Project *GetIdTokenScopeDomainOrProjectBody `json:"project,omitempty" xml:"project"`
}

func (o GetIdTokenIdScopeBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetIdTokenIdScopeBody struct{}"
	}

	return strings.Join([]string{"GetIdTokenIdScopeBody", string(data)}, " ")
}
