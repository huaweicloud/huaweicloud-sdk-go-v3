package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ApiAclCreate struct {

	// ACL策略名称。支持汉字，英文，数字，下划线，且只能以英文和汉字开头，3 ~ 64字符。 > 中文字符必须为UTF-8或者unicode编码。
	AclName string `json:"acl_name"`

	// 类型 -  PERMIT (白名单类型) -  DENY (黑名单类型)
	AclType ApiAclCreateAclType `json:"acl_type"`

	// ACL策略值，支持一个或多个值，使用英文半角逗号分隔。 - entity_type为IP时，策略值需填写IP地址，最多可支持100个IP地址。 - entity_type为DOMAIN时，策略值需填写账号名，账号支持除英文半角逗号以外的任意ASCII字符，账号名长度限制在1-64个字符，不支持纯数字。多账号名字符的总长度不超过1024。 - entity_type为DOMAIN_ID时，策略值需填写账号ID，获取方式请参见API参考的“附录 > 获取账号ID”章节。
	AclValue string `json:"acl_value"`

	// 对象类型： - IP：IP地址 - DOMAIN：账号名 - DOMAIN_ID：账号ID
	EntityType ApiAclCreateEntityType `json:"entity_type"`
}

func (o ApiAclCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiAclCreate struct{}"
	}

	return strings.Join([]string{"ApiAclCreate", string(data)}, " ")
}

type ApiAclCreateAclType struct {
	value string
}

type ApiAclCreateAclTypeEnum struct {
	PERMIT ApiAclCreateAclType
	DENY   ApiAclCreateAclType
}

func GetApiAclCreateAclTypeEnum() ApiAclCreateAclTypeEnum {
	return ApiAclCreateAclTypeEnum{
		PERMIT: ApiAclCreateAclType{
			value: "PERMIT",
		},
		DENY: ApiAclCreateAclType{
			value: "DENY",
		},
	}
}

func (c ApiAclCreateAclType) Value() string {
	return c.value
}

func (c ApiAclCreateAclType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiAclCreateAclType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ApiAclCreateEntityType struct {
	value string
}

type ApiAclCreateEntityTypeEnum struct {
	IP        ApiAclCreateEntityType
	DOMAIN    ApiAclCreateEntityType
	DOMAIN_ID ApiAclCreateEntityType
}

func GetApiAclCreateEntityTypeEnum() ApiAclCreateEntityTypeEnum {
	return ApiAclCreateEntityTypeEnum{
		IP: ApiAclCreateEntityType{
			value: "IP",
		},
		DOMAIN: ApiAclCreateEntityType{
			value: "DOMAIN",
		},
		DOMAIN_ID: ApiAclCreateEntityType{
			value: "DOMAIN_ID",
		},
	}
}

func (c ApiAclCreateEntityType) Value() string {
	return c.value
}

func (c ApiAclCreateEntityType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiAclCreateEntityType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
