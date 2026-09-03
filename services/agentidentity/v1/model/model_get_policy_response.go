package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetPolicyResponse Response Object
type GetPolicyResponse struct {
	Policy         *Policy `json:"policy,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o GetPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetPolicyResponse struct{}"
	}

	return strings.Join([]string{"GetPolicyResponse", string(data)}, " ")
}
