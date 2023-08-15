package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InviteWithPwdResponse Response Object
type InviteWithPwdResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o InviteWithPwdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InviteWithPwdResponse struct{}"
	}

	return strings.Join([]string{"InviteWithPwdResponse", string(data)}, " ")
}
