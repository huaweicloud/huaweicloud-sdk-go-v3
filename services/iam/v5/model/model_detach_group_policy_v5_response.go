package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachGroupPolicyV5Response Response Object
type DetachGroupPolicyV5Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DetachGroupPolicyV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachGroupPolicyV5Response struct{}"
	}

	return strings.Join([]string{"DetachGroupPolicyV5Response", string(data)}, " ")
}
