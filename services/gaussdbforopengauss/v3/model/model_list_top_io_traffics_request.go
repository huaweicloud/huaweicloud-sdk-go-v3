package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListTopIoTrafficsRequest Request Object
type ListTopIoTrafficsRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	// 节点ID。节点应为CN或者非日志角色的DN节点，并且节点状态为正常。
	NodeId string `json:"node_id"`

	// 组件ID。组件应为CN或者非日志角色的DN组件。 DN：Data Node，和CN对应的概念。负责实际执行表数据的存储、查询操作。 CN：Coordinator Node，负责数据库系统元数据存储、查询任务的分解和部分执行，以及将DN中查询结果汇聚在一起。
	ComponentId string `json:"component_id"`

	// 期望查询数据库进程下TOP IO线程数（默认值为20）。接口返回TOP IO线程与会话信息关联后的结果，数量最大不超过该值。
	TopIoNum *int32 `json:"top_io_num,omitempty"`

	// TOP IO排序条件
	SortCondition *ListTopIoTrafficsRequestSortCondition `json:"sort_condition,omitempty"`
}

func (o ListTopIoTrafficsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopIoTrafficsRequest struct{}"
	}

	return strings.Join([]string{"ListTopIoTrafficsRequest", string(data)}, " ")
}

type ListTopIoTrafficsRequestSortCondition struct {
	value string
}

type ListTopIoTrafficsRequestSortConditionEnum struct {
	READ  ListTopIoTrafficsRequestSortCondition
	WRITE ListTopIoTrafficsRequestSortCondition
}

func GetListTopIoTrafficsRequestSortConditionEnum() ListTopIoTrafficsRequestSortConditionEnum {
	return ListTopIoTrafficsRequestSortConditionEnum{
		READ: ListTopIoTrafficsRequestSortCondition{
			value: "read",
		},
		WRITE: ListTopIoTrafficsRequestSortCondition{
			value: "write",
		},
	}
}

func (c ListTopIoTrafficsRequestSortCondition) Value() string {
	return c.value
}

func (c ListTopIoTrafficsRequestSortCondition) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTopIoTrafficsRequestSortCondition) UnmarshalJSON(b []byte) error {
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
