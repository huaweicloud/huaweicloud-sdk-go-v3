package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAccessPolicyResponse Response Object
type CreateAccessPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateAccessPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAccessPolicyResponse struct{}"
	}

	return strings.Join([]string{"CreateAccessPolicyResponse", string(data)}, " ")
}
