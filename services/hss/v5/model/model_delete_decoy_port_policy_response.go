package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDecoyPortPolicyResponse Response Object
type DeleteDecoyPortPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDecoyPortPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDecoyPortPolicyResponse struct{}"
	}

	return strings.Join([]string{"DeleteDecoyPortPolicyResponse", string(data)}, " ")
}
