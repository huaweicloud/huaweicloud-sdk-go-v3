package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachAgencyPolicyV5Response Response Object
type AttachAgencyPolicyV5Response struct {
	HttpStatusCode int `json:"-"`
}

func (o AttachAgencyPolicyV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachAgencyPolicyV5Response struct{}"
	}

	return strings.Join([]string{"AttachAgencyPolicyV5Response", string(data)}, " ")
}
