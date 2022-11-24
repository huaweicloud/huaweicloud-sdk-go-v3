package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResetRestartRequestBodyPhones struct {

	// 云手机id
	PhoneId string `json:"phone_id"`

	// 云手机属性列表
	Property *string `json:"property,omitempty"`
}

func (o ResetRestartRequestBodyPhones) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetRestartRequestBodyPhones struct{}"
	}

	return strings.Join([]string{"ResetRestartRequestBodyPhones", string(data)}, " ")
}
