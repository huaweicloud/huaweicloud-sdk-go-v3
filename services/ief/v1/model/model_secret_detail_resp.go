package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 密钥
type SecretDetailResp struct {

	// 密钥ID
	Id string `json:"id"`

	// 密钥名称，以小写英文字母开头，4-64位，可以使用小写英文、数字、中划线（-），不能以中划线结尾
	Name string `json:"name"`

	// 密钥描述,最大长度255，不允许^ ~ # $ % & * < > ( ) [ ] { } ' \" \\
	Description string `json:"description"`

	// secrets是一个字典，由多个键值对组成，json化后最大总长度为1048576，key和value均为字符串。键值对中key由大小写字母或中划线开头，由数字、大小写字母、点号（.）、中划线（-）、下划线（_）组成，最小长度为1，最大长度63个字符, 键值对中的value必须为base64字符。 注：secrets字典的长度即字典转为标准的字符串后的长度，例如字典{\"a\": \"b\"}转为标准字符串后为'{\"a\": \"b\"}'，长度为10
	Secrets map[string]string `json:"secrets"`

	// 项目ID
	ProjectId string `json:"project_id"`

	// 创建时间
	CreatedAt string `json:"created_at"`

	// 更新时间
	UpdatedAt string `json:"updated_at"`

	// 密钥类型，目前只支持“Opaque”类型
	Type string `json:"type"`
}

func (o SecretDetailResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecretDetailResp struct{}"
	}

	return strings.Join([]string{"SecretDetailResp", string(data)}, " ")
}
