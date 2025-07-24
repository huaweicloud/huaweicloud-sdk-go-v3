package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListServersRequest Request Object
type ListServersRequest struct {

	// 服务器资产管理状态： - delivering：等待物料。 - received：到货，物料到货进场。 - onboard：上架，对物理服务器进行机柜机架、物理组网、BMC地址配置等，可通过BMC远程管理。 - ready：可用，完成网调、软调及转维验收。 - in-use：使用中，创建裸机实例。 - frozen：冻结，因欠费导致资源冻结。 - offboarding：下架中。
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
	DELIVERING  ListServersRequestManageState
	RECEIVED    ListServersRequestManageState
	ONBOARD     ListServersRequestManageState
	READY       ListServersRequestManageState
	FROZEN      ListServersRequestManageState
	OFFBOARDING ListServersRequestManageState
}

func GetListServersRequestManageStateEnum() ListServersRequestManageStateEnum {
	return ListServersRequestManageStateEnum{
		DELIVERING: ListServersRequestManageState{
			value: "delivering",
		},
		RECEIVED: ListServersRequestManageState{
			value: "received",
		},
		ONBOARD: ListServersRequestManageState{
			value: "onboard",
		},
		READY: ListServersRequestManageState{
			value: "ready",
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
