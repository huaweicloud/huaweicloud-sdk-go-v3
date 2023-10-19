package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CentralNetworkPolicyStateEnum 中心网络策略状态。 - AVAILABLE (可用) - CANCELING (取消中) - APPLYING (应用中) - FAILED (失败) - DELETED (已刪除)
type CentralNetworkPolicyStateEnum struct {
	value string
}

type CentralNetworkPolicyStateEnumEnum struct {
	AVAILABLE CentralNetworkPolicyStateEnum
	CANCELING CentralNetworkPolicyStateEnum
	APPLYING  CentralNetworkPolicyStateEnum
	FAILED    CentralNetworkPolicyStateEnum
	DELETED   CentralNetworkPolicyStateEnum
}

func GetCentralNetworkPolicyStateEnumEnum() CentralNetworkPolicyStateEnumEnum {
	return CentralNetworkPolicyStateEnumEnum{
		AVAILABLE: CentralNetworkPolicyStateEnum{
			value: "AVAILABLE",
		},
		CANCELING: CentralNetworkPolicyStateEnum{
			value: "CANCELING",
		},
		APPLYING: CentralNetworkPolicyStateEnum{
			value: "APPLYING",
		},
		FAILED: CentralNetworkPolicyStateEnum{
			value: "FAILED",
		},
		DELETED: CentralNetworkPolicyStateEnum{
			value: "DELETED",
		},
	}
}

func (c CentralNetworkPolicyStateEnum) Value() string {
	return c.value
}

func (c CentralNetworkPolicyStateEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CentralNetworkPolicyStateEnum) UnmarshalJSON(b []byte) error {
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
