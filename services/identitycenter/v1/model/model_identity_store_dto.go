package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type IdentityStoreDto struct {

	// 关联到IAM身份中心实例的身份源的全局唯一标识符（ID）。
	IdentityStoreId string `json:"identity_store_id"`

	// 身份源类型
	IdentityStoreType IdentityStoreDtoIdentityStoreType `json:"identity_store_type"`

	// 登录认证类型
	AuthenticationType IdentityStoreDtoAuthenticationType `json:"authentication_type"`

	// 预配类型
	ProvisioningType []IdentityStoreDtoProvisioningType `json:"provisioning_type"`

	// 身份源是否启用状态
	Status IdentityStoreDtoStatus `json:"status"`
}

func (o IdentityStoreDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdentityStoreDto struct{}"
	}

	return strings.Join([]string{"IdentityStoreDto", string(data)}, " ")
}

type IdentityStoreDtoIdentityStoreType struct {
	value string
}

type IdentityStoreDtoIdentityStoreTypeEnum struct {
	USER_POOL IdentityStoreDtoIdentityStoreType
}

func GetIdentityStoreDtoIdentityStoreTypeEnum() IdentityStoreDtoIdentityStoreTypeEnum {
	return IdentityStoreDtoIdentityStoreTypeEnum{
		USER_POOL: IdentityStoreDtoIdentityStoreType{
			value: "UserPool",
		},
	}
}

func (c IdentityStoreDtoIdentityStoreType) Value() string {
	return c.value
}

func (c IdentityStoreDtoIdentityStoreType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IdentityStoreDtoIdentityStoreType) UnmarshalJSON(b []byte) error {
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

type IdentityStoreDtoAuthenticationType struct {
	value string
}

type IdentityStoreDtoAuthenticationTypeEnum struct {
	SAML_2_0 IdentityStoreDtoAuthenticationType
	DEFAULT  IdentityStoreDtoAuthenticationType
}

func GetIdentityStoreDtoAuthenticationTypeEnum() IdentityStoreDtoAuthenticationTypeEnum {
	return IdentityStoreDtoAuthenticationTypeEnum{
		SAML_2_0: IdentityStoreDtoAuthenticationType{
			value: "SAML_2.0",
		},
		DEFAULT: IdentityStoreDtoAuthenticationType{
			value: "DEFAULT",
		},
	}
}

func (c IdentityStoreDtoAuthenticationType) Value() string {
	return c.value
}

func (c IdentityStoreDtoAuthenticationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IdentityStoreDtoAuthenticationType) UnmarshalJSON(b []byte) error {
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

type IdentityStoreDtoProvisioningType struct {
	value string
}

type IdentityStoreDtoProvisioningTypeEnum struct {
	DEFAULT IdentityStoreDtoProvisioningType
	SCIM    IdentityStoreDtoProvisioningType
}

func GetIdentityStoreDtoProvisioningTypeEnum() IdentityStoreDtoProvisioningTypeEnum {
	return IdentityStoreDtoProvisioningTypeEnum{
		DEFAULT: IdentityStoreDtoProvisioningType{
			value: "DEFAULT",
		},
		SCIM: IdentityStoreDtoProvisioningType{
			value: "SCIM",
		},
	}
}

func (c IdentityStoreDtoProvisioningType) Value() string {
	return c.value
}

func (c IdentityStoreDtoProvisioningType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IdentityStoreDtoProvisioningType) UnmarshalJSON(b []byte) error {
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

type IdentityStoreDtoStatus struct {
	value string
}

type IdentityStoreDtoStatusEnum struct {
	DISABLED IdentityStoreDtoStatus
	ENABLED  IdentityStoreDtoStatus
}

func GetIdentityStoreDtoStatusEnum() IdentityStoreDtoStatusEnum {
	return IdentityStoreDtoStatusEnum{
		DISABLED: IdentityStoreDtoStatus{
			value: "DISABLED",
		},
		ENABLED: IdentityStoreDtoStatus{
			value: "ENABLED",
		},
	}
}

func (c IdentityStoreDtoStatus) Value() string {
	return c.value
}

func (c IdentityStoreDtoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IdentityStoreDtoStatus) UnmarshalJSON(b []byte) error {
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
