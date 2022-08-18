package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PhonesPhones struct {

	// 云手机id
	PhoneId string `json:"phone_id"`

	// 云手机属性列表
	Property string `json:"property"`
}

func (o PhonesPhones) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PhonesPhones struct{}"
	}

	return strings.Join([]string{"PhonesPhones", string(data)}, " ")
}
