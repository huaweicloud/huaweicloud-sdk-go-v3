package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListServersRequest Request Object
type ListServersRequest struct {

	// **参数解释**： 服务器管理状态 **约束限制**： 不涉及 **取值范围**： - onboard：上架中，用户下单，完成LLD设计。 - ready：交付完成，完成硬装、网调、服务器初始化、软调及转维验收。 - in-use：使用中，用户发放裸机。 - frozen：冻结，因欠费导致资源冻结。 - offboarding：下架中。  **默认取值**： 不涉及
	ManageState *ListServersRequestManageState `json:"manage_state,omitempty"`

	// 上一页数据的最后一条记录的id
	Marker *string `json:"marker,omitempty"`

	// 分页查询时每页行数。最大值为 1000。  默认值： 当不设置值或设置的值小于 10 时，默认值为 10。 当设置的值大于 1000 时，默认值为 1000。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServersRequest struct{}"
	}

	return strings.Join([]string{"ListServersRequest", string(data)}, " ")
}

type ListServersRequestManageState struct {
	value string
}

type ListServersRequestManageStateEnum struct {
	ONBOARD     ListServersRequestManageState
	READY       ListServersRequestManageState
	IN_USE      ListServersRequestManageState
	FROZEN      ListServersRequestManageState
	OFFBOARDING ListServersRequestManageState
}

func GetListServersRequestManageStateEnum() ListServersRequestManageStateEnum {
	return ListServersRequestManageStateEnum{
		ONBOARD: ListServersRequestManageState{
			value: "onboard",
		},
		READY: ListServersRequestManageState{
			value: "ready",
		},
		IN_USE: ListServersRequestManageState{
			value: "in-use",
		},
		FROZEN: ListServersRequestManageState{
			value: "frozen",
		},
		OFFBOARDING: ListServersRequestManageState{
			value: "offboarding",
		},
	}
}

func (c ListServersRequestManageState) Value() string {
	return c.value
}

func (c ListServersRequestManageState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListServersRequestManageState) UnmarshalJSON(b []byte) error {
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
