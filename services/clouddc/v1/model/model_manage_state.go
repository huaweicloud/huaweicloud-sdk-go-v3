package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ManageState **参数解释**： 服务器管理状态 **约束限制**： 不涉及 **取值范围** - onboard：上架中，用户下单，完成LLD设计。 - ready：交付完成，完成硬装、网调、服务器初始化、软调及转维验收。 - in-use：使用中，用户发放裸机。 - frozen：冻结，因欠费导致资源冻结。 - offboarding：下架中。  ```mermaid stateDiagram-v2    [*] --> onboard : 完成LLD设计   onboard --> ready : 完成网调、服务器初始化、软调及转维验收   ready --> in_use : 发放裸机实例   ready --> offboarding : 请求下架   ready --> frozen : 欠费      in_use --> ready : 删除裸机实例   in_use --> frozen : 欠费    frozen --> offboarding : 请求下架   in_use --> offboarding : 请求下架   offboarding --> [*] : 完成下架   state \"in-use\" as in_use ```
type ManageState struct {
	value string
}

type ManageStateEnum struct {
	DELIVERING ManageState
	RECEIVED   ManageState
	ONBOARD    ManageState
	READY      ManageState
	FROZEN     ManageState
}

func GetManageStateEnum() ManageStateEnum {
	return ManageStateEnum{
		DELIVERING: ManageState{
			value: "delivering",
		},
		RECEIVED: ManageState{
			value: "received",
		},
		ONBOARD: ManageState{
			value: "onboard",
		},
		READY: ManageState{
			value: "ready",
		},
		FROZEN: ManageState{
			value: "frozen",
		},
	}
}

func (c ManageState) Value() string {
	return c.value
}

func (c ManageState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ManageState) UnmarshalJSON(b []byte) error {
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
