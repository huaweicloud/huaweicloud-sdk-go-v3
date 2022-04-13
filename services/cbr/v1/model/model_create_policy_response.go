package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreatePolicyResponse struct {
	Policy         *Policy `json:"policy,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePolicyResponse struct{}"
	}

	return strings.Join([]string{"CreatePolicyResponse", string(data)}, " ")
}
