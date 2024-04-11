package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResponseInfoHeader struct {

	// 分组自定义响应的响应头的key，支持英文字母、数字和中划线，长度为1到128位
	Key *string `json:"key,omitempty"`

	// 分组自定义响应的响应头的value，为长度为1到1024位的字符串
	Value *string `json:"value,omitempty"`
}

func (o ResponseInfoHeader) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResponseInfoHeader struct{}"
	}

	return strings.Join([]string{"ResponseInfoHeader", string(data)}, " ")
}
