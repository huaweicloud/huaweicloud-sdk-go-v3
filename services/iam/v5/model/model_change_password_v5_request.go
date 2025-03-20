package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangePasswordV5Request Request Object
type ChangePasswordV5Request struct {
	Body *ChangePasswordReqBody `json:"body,omitempty"`
}

func (o ChangePasswordV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangePasswordV5Request struct{}"
	}

	return strings.Join([]string{"ChangePasswordV5Request", string(data)}, " ")
}
