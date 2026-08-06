package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteKeyPolicyResponse Response Object
type DeleteKeyPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteKeyPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteKeyPolicyResponse struct{}"
	}

	return strings.Join([]string{"DeleteKeyPolicyResponse", string(data)}, " ")
}
