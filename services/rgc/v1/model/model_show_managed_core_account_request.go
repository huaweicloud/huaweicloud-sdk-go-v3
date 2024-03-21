package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowManagedCoreAccountRequest Request Object
type ShowManagedCoreAccountRequest struct {

	// 纳管账号类型。类型包括LOGGING，SECURITY和PRIMARY。
	AccountType ShowManagedCoreAccountRequestAccountType `json:"account_type"`
}

func (o ShowManagedCoreAccountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowManagedCoreAccountRequest struct{}"
	}

	return strings.Join([]string{"ShowManagedCoreAccountRequest", string(data)}, " ")
}

type ShowManagedCoreAccountRequestAccountType struct {
	value string
}

type ShowManagedCoreAccountRequestAccountTypeEnum struct {
	LOGGING  ShowManagedCoreAccountRequestAccountType
	SECURITY ShowManagedCoreAccountRequestAccountType
	PRIMARY  ShowManagedCoreAccountRequestAccountType
}

func GetShowManagedCoreAccountRequestAccountTypeEnum() ShowManagedCoreAccountRequestAccountTypeEnum {
	return ShowManagedCoreAccountRequestAccountTypeEnum{
		LOGGING: ShowManagedCoreAccountRequestAccountType{
			value: "LOGGING",
		},
		SECURITY: ShowManagedCoreAccountRequestAccountType{
			value: "SECURITY",
		},
		PRIMARY: ShowManagedCoreAccountRequestAccountType{
			value: "PRIMARY",
		},
	}
}

func (c ShowManagedCoreAccountRequestAccountType) Value() string {
	return c.value
}

func (c ShowManagedCoreAccountRequestAccountType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowManagedCoreAccountRequestAccountType) UnmarshalJSON(b []byte) error {
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
