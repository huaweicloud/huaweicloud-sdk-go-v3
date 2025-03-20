package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachAgencyPolicyV5Response Response Object
type DetachAgencyPolicyV5Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DetachAgencyPolicyV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachAgencyPolicyV5Response struct{}"
	}

	return strings.Join([]string{"DetachAgencyPolicyV5Response", string(data)}, " ")
}
