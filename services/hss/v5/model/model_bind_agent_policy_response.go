package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BindAgentPolicyResponse Response Object
type BindAgentPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BindAgentPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindAgentPolicyResponse struct{}"
	}

	return strings.Join([]string{"BindAgentPolicyResponse", string(data)}, " ")
}
