package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePasswordPolicyResponse Response Object
type UpdatePasswordPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdatePasswordPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePasswordPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdatePasswordPolicyResponse", string(data)}, " ")
}
