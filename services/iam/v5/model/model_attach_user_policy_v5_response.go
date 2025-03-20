package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachUserPolicyV5Response Response Object
type AttachUserPolicyV5Response struct {
	HttpStatusCode int `json:"-"`
}

func (o AttachUserPolicyV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachUserPolicyV5Response struct{}"
	}

	return strings.Join([]string{"AttachUserPolicyV5Response", string(data)}, " ")
}
