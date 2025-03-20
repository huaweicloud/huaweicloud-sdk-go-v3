package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetPolicyV5Response Response Object
type GetPolicyV5Response struct {
	Policy         *Policy `json:"policy,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o GetPolicyV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetPolicyV5Response struct{}"
	}

	return strings.Join([]string{"GetPolicyV5Response", string(data)}, " ")
}
