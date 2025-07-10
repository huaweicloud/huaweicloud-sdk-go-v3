package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AccountBaselineRsp 纳管账号基本信息。
type AccountBaselineRsp struct {

	// 纳管账号名称。
	AccountName string `json:"account_name"`

	// 纳管账号的唯一标识符（ID）。
	AccountId *string `json:"account_id,omitempty"`

	// 纳管账号类型。类型包括LOGGING，SECURITY。
	AccountType AccountBaselineRspAccountType `json:"account_type"`
}

func (o AccountBaselineRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccountBaselineRsp struct{}"
	}

	return strings.Join([]string{"AccountBaselineRsp", string(data)}, " ")
}

type AccountBaselineRspAccountType struct {
	value string
}

type AccountBaselineRspAccountTypeEnum struct {
	LOGGING  AccountBaselineRspAccountType
	SECURITY AccountBaselineRspAccountType
}

func GetAccountBaselineRspAccountTypeEnum() AccountBaselineRspAccountTypeEnum {
	return AccountBaselineRspAccountTypeEnum{
		LOGGING: AccountBaselineRspAccountType{
			value: "LOGGING",
		},
		SECURITY: AccountBaselineRspAccountType{
			value: "SECURITY",
		},
	}
}

func (c AccountBaselineRspAccountType) Value() string {
	return c.value
}

func (c AccountBaselineRspAccountType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccountBaselineRspAccountType) UnmarshalJSON(b []byte) error {
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
