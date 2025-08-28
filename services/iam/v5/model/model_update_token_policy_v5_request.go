package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTokenPolicyV5Request Request Object
type UpdateTokenPolicyV5Request struct {
	Body *UpdateTokenPolicyReqBody `json:"body,omitempty"`
}

func (o UpdateTokenPolicyV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTokenPolicyV5Request struct{}"
	}

	return strings.Join([]string{"UpdateTokenPolicyV5Request", string(data)}, " ")
}
