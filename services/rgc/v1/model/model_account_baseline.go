package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AccountBaseline 账号基本信息。
type AccountBaseline struct {

	// 账号名称。
	AccountName string `json:"account_name"`

	// 手机号码。
	Phone *string `json:"phone,omitempty"`

	// 账号邮箱。
	AccountEmail string `json:"account_email"`

	// 账号类型logging,security。 * LOGGING - 日志账号 * SECURITY - 安全账号 * CUSTOM - 自定义账号
	AccountType AccountBaselineAccountType `json:"account_type"`
}

func (o AccountBaseline) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccountBaseline struct{}"
	}

	return strings.Join([]string{"AccountBaseline", string(data)}, " ")
}

type AccountBaselineAccountType struct {
	value string
}

type AccountBaselineAccountTypeEnum struct {
	LOGGING  AccountBaselineAccountType
	SECURITY AccountBaselineAccountType
	CUSTOM   AccountBaselineAccountType
}

func GetAccountBaselineAccountTypeEnum() AccountBaselineAccountTypeEnum {
	return AccountBaselineAccountTypeEnum{
		LOGGING: AccountBaselineAccountType{
			value: "LOGGING",
		},
		SECURITY: AccountBaselineAccountType{
			value: "SECURITY",
		},
		CUSTOM: AccountBaselineAccountType{
			value: "CUSTOM",
		},
	}
}

func (c AccountBaselineAccountType) Value() string {
	return c.value
}

func (c AccountBaselineAccountType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccountBaselineAccountType) UnmarshalJSON(b []byte) error {
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
