package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeploymentSecrets struct {

	// 密钥的名称，以英文小写字母开头，由中文字符，英文字母，数字，下划线和中划线组成，不能以中划线结尾，长度4-64位
	Name *string `json:"name,omitempty"`

	// 密钥的属性名，以英文字母和中划线开头，由英文字母、数字、点号、中划线和下划线组成，长度1-63位
	Key *string `json:"key,omitempty"`
}

func (o DeploymentSecrets) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeploymentSecrets struct{}"
	}

	return strings.Join([]string{"DeploymentSecrets", string(data)}, " ")
}
