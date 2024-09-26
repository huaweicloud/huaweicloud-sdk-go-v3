package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CentralNetworkStateEnum 中心网络状态。 - AVAILABLE (可用) - UPDATING (处理中) - FAILED (失败) - CREATING (创建中) - DELETING (删除中) - DELETED (已刪除) - RESTORING (恢复中)
type CentralNetworkStateEnum struct {
	value string
}

type CentralNetworkStateEnumEnum struct {
	AVAILABLE CentralNetworkStateEnum
	UPDATING  CentralNetworkStateEnum
	FAILED    CentralNetworkStateEnum
	CREATING  CentralNetworkStateEnum
	DELETING  CentralNetworkStateEnum
	DELETED   CentralNetworkStateEnum
	RESTORING CentralNetworkStateEnum
}

func GetCentralNetworkStateEnumEnum() CentralNetworkStateEnumEnum {
	return CentralNetworkStateEnumEnum{
		AVAILABLE: CentralNetworkStateEnum{
			value: "AVAILABLE",
		},
		UPDATING: CentralNetworkStateEnum{
			value: "UPDATING",
		},
		FAILED: CentralNetworkStateEnum{
			value: "FAILED",
		},
		CREATING: CentralNetworkStateEnum{
			value: "CREATING",
		},
		DELETING: CentralNetworkStateEnum{
			value: "DELETING",
		},
		DELETED: CentralNetworkStateEnum{
			value: "DELETED",
		},
		RESTORING: CentralNetworkStateEnum{
			value: "RESTORING",
		},
	}
}

func (c CentralNetworkStateEnum) Value() string {
	return c.value
}

func (c CentralNetworkStateEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CentralNetworkStateEnum) UnmarshalJSON(b []byte) error {
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
