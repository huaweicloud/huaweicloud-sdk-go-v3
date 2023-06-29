package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecretObjectInSecretRequestBody 创建密钥或者更更新密钥请求体中的secret对象
type SecretObjectInSecretRequestBody struct {

	// 密钥名称，以小写英文字母开头，4-64位，可以使用小写英文、数字、中划线（-），不能以中划线结尾。
	Name string `json:"name"`

	// 密钥描述,最大长度255，不允许^ ~ # $ % & * < > ( ) [ ] { } ' \" \\。
	Description *string `json:"description,omitempty"`

	// secrets是一个列表，由多个字典组成，数量上限为100个。字典的key和value均为字符串。key由大小写字母，中划线（`-`）开头，由数字、大小写字母、点号（`.`）、中划线（`-`）、下划线（`_`）组成，最小长度为1，最大长度63个字符, value不能含有字符不允许^ ~ # $ \\|% & < > ( ) [ ] { } ' \" \\，最大长度10000个字符。
	Secrets []Secret `json:"secrets"`

	// tags是一个列表，由多个字典组成，数量上限为20个。字典的key和value均为字符串。key由数字、大小写字母、点号（`.`）、中划线（`-`）、下划线（`_`）组成，最小长度为1，最大长度36个字符, value由数字、大小写字母、点号（`.`）、中划线（`-`）、下划线（`_`）, 斜线（`\\`）组成，最小长度为0，最大长度43个字符。
	Tags []Tag `json:"tags"`
}

func (o SecretObjectInSecretRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecretObjectInSecretRequestBody struct{}"
	}

	return strings.Join([]string{"SecretObjectInSecretRequestBody", string(data)}, " ")
}
