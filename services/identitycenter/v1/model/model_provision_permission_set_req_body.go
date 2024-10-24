package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ProvisionPermissionSetReqBody ProvisionPermissionSet的请求体
type ProvisionPermissionSetReqBody struct {

	// 账号ID
	TargetId *string `json:"target_id,omitempty"`

	// 创建绑定的实体类型
	TargetType ProvisionPermissionSetReqBodyTargetType `json:"target_type"`
}

func (o ProvisionPermissionSetReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProvisionPermissionSetReqBody struct{}"
	}

	return strings.Join([]string{"ProvisionPermissionSetReqBody", string(data)}, " ")
}

type ProvisionPermissionSetReqBodyTargetType struct {
	value string
}

type ProvisionPermissionSetReqBodyTargetTypeEnum struct {
	ACCOUNT                  ProvisionPermissionSetReqBodyTargetType
	ALL_PROVISIONED_ACCOUNTS ProvisionPermissionSetReqBodyTargetType
}

func GetProvisionPermissionSetReqBodyTargetTypeEnum() ProvisionPermissionSetReqBodyTargetTypeEnum {
	return ProvisionPermissionSetReqBodyTargetTypeEnum{
		ACCOUNT: ProvisionPermissionSetReqBodyTargetType{
			value: "ACCOUNT",
		},
		ALL_PROVISIONED_ACCOUNTS: ProvisionPermissionSetReqBodyTargetType{
			value: "ALL_PROVISIONED_ACCOUNTS",
		},
	}
}

func (c ProvisionPermissionSetReqBodyTargetType) Value() string {
	return c.value
}

func (c ProvisionPermissionSetReqBodyTargetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ProvisionPermissionSetReqBodyTargetType) UnmarshalJSON(b []byte) error {
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
