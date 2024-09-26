package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CentralNetworkConnectionStateEnum 中心网络连接状态。 - AVAILABLE (可用) - CREATING (创建中) - UPDATING (更新中) - DELETING (删除中) - FREEZING (冻结中) - UNFREEZING (解冻中) - RECOVERING (恢复中) - FAILED (失败) - DELETED (已刪除) - APPROVING (审批中) - APPROVED (已审批) - UNAPPROVED (审批未通过)
type CentralNetworkConnectionStateEnum struct {
	value string
}

type CentralNetworkConnectionStateEnumEnum struct {
	AVAILABLE  CentralNetworkConnectionStateEnum
	CREATING   CentralNetworkConnectionStateEnum
	UPDATING   CentralNetworkConnectionStateEnum
	DELETING   CentralNetworkConnectionStateEnum
	FREEZING   CentralNetworkConnectionStateEnum
	UNFREEZING CentralNetworkConnectionStateEnum
	RECOVERING CentralNetworkConnectionStateEnum
	FAILED     CentralNetworkConnectionStateEnum
	DELETED    CentralNetworkConnectionStateEnum
	APPROVING  CentralNetworkConnectionStateEnum
	APPROVED   CentralNetworkConnectionStateEnum
	UNAPPROVED CentralNetworkConnectionStateEnum
}

func GetCentralNetworkConnectionStateEnumEnum() CentralNetworkConnectionStateEnumEnum {
	return CentralNetworkConnectionStateEnumEnum{
		AVAILABLE: CentralNetworkConnectionStateEnum{
			value: "AVAILABLE",
		},
		CREATING: CentralNetworkConnectionStateEnum{
			value: "CREATING",
		},
		UPDATING: CentralNetworkConnectionStateEnum{
			value: "UPDATING",
		},
		DELETING: CentralNetworkConnectionStateEnum{
			value: "DELETING",
		},
		FREEZING: CentralNetworkConnectionStateEnum{
			value: "FREEZING",
		},
		UNFREEZING: CentralNetworkConnectionStateEnum{
			value: "UNFREEZING",
		},
		RECOVERING: CentralNetworkConnectionStateEnum{
			value: "RECOVERING",
		},
		FAILED: CentralNetworkConnectionStateEnum{
			value: "FAILED",
		},
		DELETED: CentralNetworkConnectionStateEnum{
			value: "DELETED",
		},
		APPROVING: CentralNetworkConnectionStateEnum{
			value: "APPROVING",
		},
		APPROVED: CentralNetworkConnectionStateEnum{
			value: "APPROVED",
		},
		UNAPPROVED: CentralNetworkConnectionStateEnum{
			value: "UNAPPROVED",
		},
	}
}

func (c CentralNetworkConnectionStateEnum) Value() string {
	return c.value
}

func (c CentralNetworkConnectionStateEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CentralNetworkConnectionStateEnum) UnmarshalJSON(b []byte) error {
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
