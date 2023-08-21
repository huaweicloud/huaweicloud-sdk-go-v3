package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListEdgeNodesRequest Request Object
type ListEdgeNodesRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	// 边缘节点名称，模糊匹配
	Name *string `json:"name,omitempty"`

	// 每页显示的条目数量，取值范围1~1000，默认为500
	Limit *string `json:"limit,omitempty"`

	// 查询的起始位置，取值范围为非负整数，默认为0
	Offset *string `json:"offset,omitempty"`

	// 显示的条目排列顺序，使用:分隔参考值和顺序， 如sort=created_at%3Adesc表示根据创建时间逆序排列
	Sort *string `json:"sort,omitempty"`

	// 按终端设备ID查找
	DeviceId *string `json:"device_id,omitempty"`

	// 按绑定终端设备名称查找
	DeviceName *string `json:"device_name,omitempty"`

	// 按应用名称查找
	AppName *string `json:"app_name,omitempty"`

	// 按边缘节点状态查找，节点状态可选项： - RUNNING：运行中 - FAIL：故障 - UPGRADING：升级中 - STOPPED：已停用 - UNCONNECTED：未纳管
	State *ListEdgeNodesRequestState `json:"state,omitempty"`

	// 标签的key和value通过点连接， 多个标签通过逗号连接，如：tags=key1.value1,key2.value2
	Tags *string `json:"tags,omitempty"`

	// 按边缘节点组ID查找。仅支持在铂金版实例中使用
	GroupId *string `json:"group_id,omitempty"`
}

func (o ListEdgeNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEdgeNodesRequest struct{}"
	}

	return strings.Join([]string{"ListEdgeNodesRequest", string(data)}, " ")
}

type ListEdgeNodesRequestState struct {
	value string
}

type ListEdgeNodesRequestStateEnum struct {
	RUNNING     ListEdgeNodesRequestState
	FAIL        ListEdgeNodesRequestState
	FREEZE      ListEdgeNodesRequestState
	UPGRADING   ListEdgeNodesRequestState
	STOPPED     ListEdgeNodesRequestState
	UNCONNECTED ListEdgeNodesRequestState
}

func GetListEdgeNodesRequestStateEnum() ListEdgeNodesRequestStateEnum {
	return ListEdgeNodesRequestStateEnum{
		RUNNING: ListEdgeNodesRequestState{
			value: "RUNNING",
		},
		FAIL: ListEdgeNodesRequestState{
			value: "FAIL",
		},
		FREEZE: ListEdgeNodesRequestState{
			value: "FREEZE",
		},
		UPGRADING: ListEdgeNodesRequestState{
			value: "UPGRADING",
		},
		STOPPED: ListEdgeNodesRequestState{
			value: "STOPPED",
		},
		UNCONNECTED: ListEdgeNodesRequestState{
			value: "UNCONNECTED",
		},
	}
}

func (c ListEdgeNodesRequestState) Value() string {
	return c.value
}

func (c ListEdgeNodesRequestState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEdgeNodesRequestState) UnmarshalJSON(b []byte) error {
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
