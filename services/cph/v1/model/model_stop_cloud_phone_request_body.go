package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopCloudPhoneRequestBody 关闭云手机请求体。
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
