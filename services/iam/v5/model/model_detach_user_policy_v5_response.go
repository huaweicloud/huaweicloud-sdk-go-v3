package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachUserPolicyV5Response Response Object
type DetachUserPolicyV5Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DetachUserPolicyV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachUserPolicyV5Response struct{}"
	}

	return strings.Join([]string{"DetachUserPolicyV5Response", string(data)}, " ")
}
