package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachGroupPolicyV5Response Response Object
type AttachGroupPolicyV5Response struct {
	HttpStatusCode int `json:"-"`
}

func (o AttachGroupPolicyV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachGroupPolicyV5Response struct{}"
	}

	return strings.Join([]string{"AttachGroupPolicyV5Response", string(data)}, " ")
}
