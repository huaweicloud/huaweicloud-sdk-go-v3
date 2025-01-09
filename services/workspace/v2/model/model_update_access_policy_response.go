package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAccessPolicyResponse Response Object
type UpdateAccessPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateAccessPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAccessPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateAccessPolicyResponse", string(data)}, " ")
}
