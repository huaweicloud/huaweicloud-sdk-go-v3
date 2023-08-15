package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomizedExtension struct {

	// 对象标识符。 > 本参数需要满足ASN1规范的点分十进制符号格式的字符串，如1.3.6.1.4.1.2011.4.99。
	ObjectIdentifier *string `json:"object_identifier,omitempty"`

	// 自定义属性内容。
	Value *string `json:"value,omitempty"`
}

func (o CustomizedExtension) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomizedExtension struct{}"
	}

	return strings.Join([]string{"CustomizedExtension", string(data)}, " ")
}
