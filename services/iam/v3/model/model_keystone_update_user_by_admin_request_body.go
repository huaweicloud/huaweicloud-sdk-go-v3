package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneUpdateUserByAdminRequestBody
type KeystoneUpdateUserByAdminRequestBody struct {
	User *KeystoneUpdateUserOption `json:"user"`
}

func (o KeystoneUpdateUserByAdminRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneUpdateUserByAdminRequestBody struct{}"
	}

	return strings.Join([]string{"KeystoneUpdateUserByAdminRequestBody", string(data)}, " ")
}
