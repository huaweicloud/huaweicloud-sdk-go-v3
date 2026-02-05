package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AppCreate struct {

	// APP的名称。支持汉字，英文，数字，下划线，且只能以英文和汉字开头，3 ~ 64个字符。 > 中文字符必须为UTF-8或者unicode编码。
	Name string `json:"name"`

	// APP描述。字符长度不能大于255。 > 中文字符必须为UTF-8或者unicode编码。
	Remark *string `json:"remark,omitempty"`

	// APP凭据的key。支持英文，数字，“_”，“-”，且只能以英文或数字开头，8 ~ 200个字符。
	AppKey *string `json:"app_key,omitempty"`

	// 密钥。支持英文，数字，“_”，“-”，“!”，“@”，“#”，“$”，“%”，且只能以英文或数字开头，8 ~ 128个字符。
	AppSecret *string `json:"app_secret,omitempty"`

	// 凭据关联的账号ID。
	RelatedDomainId *string `json:"related_domain_id,omitempty"`

	// 凭据关联的项目ID。
	RelatedProjectId *string `json:"related_project_id,omitempty"`
}

func (o AppCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppCreate struct{}"
	}

	return strings.Join([]string{"AppCreate", string(data)}, " ")
}
