package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePolicyV5Request Request Object
type CreatePolicyV5Request struct {
	Body *CreatePolicyReqBody `json:"body,omitempty"`
}

func (o CreatePolicyV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePolicyV5Request struct{}"
	}

	return strings.Join([]string{"CreatePolicyV5Request", string(data)}, " ")
}
