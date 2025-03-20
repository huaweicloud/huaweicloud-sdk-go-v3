package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePolicyV5Response Response Object
type CreatePolicyV5Response struct {
	Policy         *Policy `json:"policy,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePolicyV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePolicyV5Response struct{}"
	}

	return strings.Join([]string{"CreatePolicyV5Response", string(data)}, " ")
}
