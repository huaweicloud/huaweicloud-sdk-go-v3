package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PasswordRequest struct {

	// https密码
	Pwd string `json:"pwd" xml:"pwd"`
}

func (o PasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PasswordRequest struct{}"
	}

	return strings.Join([]string{"PasswordRequest", string(data)}, " ")
}
