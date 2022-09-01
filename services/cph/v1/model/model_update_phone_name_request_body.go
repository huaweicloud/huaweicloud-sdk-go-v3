package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePhoneNameRequestBody struct {

	// 云手机名称，必须为小写字母（a-z）、大写字母（A-Z）、数字（0-9）、中文字符、中划线-、下划线_，且不得超过60个字符。
	PhoneName string `json:"phone_name" xml:"phone_name"`
}

func (o UpdatePhoneNameRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePhoneNameRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePhoneNameRequestBody", string(data)}, " ")
}
