package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// State 工作流中定义的每一个节点。
type State struct {

	// 标识开始的state，图中必须且只有一个start为true的state。
	Start *bool `json:"start,omitempty"`

	// 必须有TYPE，TYPE值必须是3种State（DELAY，OPERATION，END）中一种。
	Type StateType `json:"type"`

	// 过滤输入参数，默认值是\"$\"，表示不过滤。必须是合法的JSONPath格式。 说明 END State没有payload_filter_in属性。
	PayloadFilterIn *string `json:"payload_filter_in,omitempty"`

	// 过滤state的输出结果，默认值是\"$\"，表示不过滤。 必须是合法的JSONPath格式。 说明 END State没有payload_filter_out属性。
	PayloadFilterOut *string `json:"payload_filter_out,omitempty"`

	// state的名字定义。 由小写字母、数字和中划线“-”组成，长度为[1, 20]。
	StateName string `json:"state_name"`

	// Action执行模式，支持串行，并行两种模式，默认串行  最小长度：1  最大长度：32  枚举值：  sequential  parallel
	ActionMode *string `json:"action_mode,omitempty"`

	// 节点中要执行的操作列表
	Actions *[]Action `json:"actions,omitempty"`

	// 创建工作流指定的下一个节点名称
	NextState *string `json:"next_state,omitempty"`

	// 当节点类型为事件延迟时填入需要延迟的时间，单位为秒
	TimeDelay *int32 `json:"time_delay,omitempty"`
}

func (o State) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "State struct{}"
	}

	return strings.Join([]string{"State", string(data)}, " ")
}

type StateType struct {
	value string
}

type StateTypeEnum struct {
	DELAY     StateType
	OPERATION StateType
	END       StateType
}

func GetStateTypeEnum() StateTypeEnum {
	return StateTypeEnum{
		DELAY: StateType{
			value: "DELAY",
		},
		OPERATION: StateType{
			value: "OPERATION",
		},
		END: StateType{
			value: "END",
		},
	}
}

func (c StateType) Value() string {
	return c.value
}

func (c StateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StateType) UnmarshalJSON(b []byte) error {
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
