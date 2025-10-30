package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyAutoNodeExpansionPolicyResponse Response Object
type ModifyAutoNodeExpansionPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ModifyAutoNodeExpansionPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyAutoNodeExpansionPolicyResponse struct{}"
	}

	return strings.Join([]string{"ModifyAutoNodeExpansionPolicyResponse", string(data)}, " ")
}
