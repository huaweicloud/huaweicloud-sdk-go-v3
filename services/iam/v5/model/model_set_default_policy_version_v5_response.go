package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetDefaultPolicyVersionV5Response Response Object
type SetDefaultPolicyVersionV5Response struct {
	HttpStatusCode int `json:"-"`
}

func (o SetDefaultPolicyVersionV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetDefaultPolicyVersionV5Response struct{}"
	}

	return strings.Join([]string{"SetDefaultPolicyVersionV5Response", string(data)}, " ")
}
