package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ApiAclCreate struct {

	// ACL策略名称。支持汉字，英文，数字，下划线，且只能以英文和汉字开头，3 ~ 64字符。 > 中文字符必须为UTF-8或者unicode编码。
	AclName string `json:"acl_name"`

	// 类型 -  PERMIT (白名单类型) -  DENY (黑名单类型)
	AclType ApiAclCreateAclType `json:"acl_type"`

	// ACL策略值，支持一个或多个值，使用英文半角逗号分隔
	AclValue string `json:"acl_value"`

	// 对象类型： - IP：IP地址 - DOMAIN：帐号名 - DOMAIN_ID：帐号ID
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
