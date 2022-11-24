package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StopCloudPhoneRequestBody struct {

	// 云手机id列表
	PhoneIds []string `json:"phone_ids"`
}

func (o StopCloudPhoneRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopCloudPhoneRequestBody struct{}"
	}

	return strings.Join([]string{"StopCloudPhoneRequestBody", string(data)}, " ")
}
