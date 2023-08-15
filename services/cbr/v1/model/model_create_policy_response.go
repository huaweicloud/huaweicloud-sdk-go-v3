package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePolicyResponse Response Object
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
