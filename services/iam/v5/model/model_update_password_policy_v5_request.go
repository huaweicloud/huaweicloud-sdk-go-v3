package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePasswordPolicyV5Request Request Object
type UpdatePasswordPolicyV5Request struct {
	Body *UpdatePasswordPolicyReqBody `json:"body,omitempty"`
}

func (o UpdatePasswordPolicyV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePasswordPolicyV5Request struct{}"
	}

	return strings.Join([]string{"UpdatePasswordPolicyV5Request", string(data)}, " ")
}
