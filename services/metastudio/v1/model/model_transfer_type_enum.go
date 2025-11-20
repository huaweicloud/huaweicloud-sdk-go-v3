package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TransferTypeEnum **参数解释**： 转移类型。默认值是TRANSFER_OUT。 **约束限制**： * 只有管理员或者开了资产转移白名单租户才有权限转出资产。 * 普通租户有权限转回已接收成功的资产，转回给转移发起方。 **取值范围**： * TRANSFER_OUT: 资产转出 * TRANSFER_BACK：资产转回
type TransferTypeEnum struct {
	value string
}

type TransferTypeEnumEnum struct {
	TRANSFER_OUT  TransferTypeEnum
	TRANSFER_BACK TransferTypeEnum
}

func GetTransferTypeEnumEnum() TransferTypeEnumEnum {
	return TransferTypeEnumEnum{
		TRANSFER_OUT: TransferTypeEnum{
			value: "TRANSFER_OUT",
		},
		TRANSFER_BACK: TransferTypeEnum{
			value: "TRANSFER_BACK",
		},
	}
}

func (c TransferTypeEnum) Value() string {
	return c.value
}

func (c TransferTypeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TransferTypeEnum) UnmarshalJSON(b []byte) error {
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
