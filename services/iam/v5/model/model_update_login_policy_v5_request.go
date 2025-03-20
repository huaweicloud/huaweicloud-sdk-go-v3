package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLoginPolicyV5Request Request Object
type UpdateLoginPolicyV5Request struct {
	Body *UpdateLoginPolicyReqBody `json:"body,omitempty"`
}

func (o UpdateLoginPolicyV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLoginPolicyV5Request struct{}"
	}

	return strings.Join([]string{"UpdateLoginPolicyV5Request", string(data)}, " ")
}
