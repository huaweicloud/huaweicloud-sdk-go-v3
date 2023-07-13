package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MicroserviceLabel struct {

	// 标签名称。  以字母或者数字开头和结尾，由字母、数字、连接符('-')、下划线('_')、点号('.')组成且63个字符之内。
	LabelName string `json:"label_name"`

	// 标签值。  以字母或者数字开头和结尾，由字母、数字、连接符('-')、下划线('_')、点号('.')组成且63个字符之内。
	LabelValue string `json:"label_value"`
}

func (o MicroserviceLabel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MicroserviceLabel struct{}"
	}

	return strings.Join([]string{"MicroserviceLabel", string(data)}, " ")
}
