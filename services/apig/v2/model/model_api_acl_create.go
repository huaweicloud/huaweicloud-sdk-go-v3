package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApiAclCreate struct {
	// ACL策略名称。支持汉字，英文，数字，下划线，且只能以英文和汉字开头，3 ~ 64字符。 > 中文字符必须为UTF-8或者unicode编码。

	AclName string `json:"acl_name"`
	// 类型 -  PERMIT (白名单类型) -  DENY (黑名单类型)

	AclType string `json:"acl_type"`
	// ACL策略值，支持一个或多个值，使用英文半角逗号分隔

	AclValue string `json:"acl_value"`
	// 对象类型： - IP - DOMAIN

	EntityType string `json:"entity_type"`
}

func (o ApiAclCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiAclCreate struct{}"
	}

	return strings.Join([]string{"ApiAclCreate", string(data)}, " ")
}
