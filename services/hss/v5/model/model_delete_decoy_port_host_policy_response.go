package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDecoyPortHostPolicyResponse Response Object
type DeleteDecoyPortHostPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDecoyPortHostPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDecoyPortHostPolicyResponse struct{}"
	}

	return strings.Join([]string{"DeleteDecoyPortHostPolicyResponse", string(data)}, " ")
}
