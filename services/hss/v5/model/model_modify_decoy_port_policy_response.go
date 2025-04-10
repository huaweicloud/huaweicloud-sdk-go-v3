package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyDecoyPortPolicyResponse Response Object
type ModifyDecoyPortPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ModifyDecoyPortPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyDecoyPortPolicyResponse struct{}"
	}

	return strings.Join([]string{"ModifyDecoyPortPolicyResponse", string(data)}, " ")
}
