package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdatePolicyResponse struct {
	Policy         *Policy `json:"policy,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdatePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdatePolicyResponse", string(data)}, " ")
}
