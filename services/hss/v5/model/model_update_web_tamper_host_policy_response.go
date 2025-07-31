package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWebTamperHostPolicyResponse Response Object
type UpdateWebTamperHostPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateWebTamperHostPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWebTamperHostPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateWebTamperHostPolicyResponse", string(data)}, " ")
}
