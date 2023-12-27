package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Event 实时作业节点事件触发配置
type Event struct {

	// 事件类型。 - KAFKA: 选择对应的连接名称与topic，当有新的kafka消息时将会触发作业运行一次 - DIS: 当前只支持监听DIS通道的新上报数据事件，每上报一条数据，触发作业运行一次。 - OBS: 选择要监听的OBS路径，如果该路径下有新增文件，则触发调度；新增的文件的路径名，可以通过变量Job.trigger.obsNewFiles引用。前提条件：该OBS路径已经配置DIS消息通知。
	EventType EventEventType `json:"event_type"`

	// DIS通道名称。通过DIS管理控制台获取通道名称：登录管理控制台。单击“数据接入服务”，左侧列表选择“通道管理”。通道管理页面中列出了用户拥有的通道
	Channel string `json:"channel"`

	// 执行失败处理策略。 - SUSPEND: 挂起 - IGNORE：忽略失败，读取下一事件
	FailPolicy *EventFailPolicy `json:"fail_policy,omitempty"`

	// 调度并发数
	Concurrent *int32 `json:"concurrent,omitempty"`

	// 读取策略。 - LAST: 从上次位置读取 - NEW：从最新位置读取
	ReadPolicy *EventReadPolicy `json:"read_policy,omitempty"`
}

func (o Event) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Event struct{}"
	}

	return strings.Join([]string{"Event", string(data)}, " ")
}

type EventEventType struct {
	value string
}

type EventEventTypeEnum struct {
	KAFKA EventEventType
	DIS   EventEventType
	OBS   EventEventType
}

func GetEventEventTypeEnum() EventEventTypeEnum {
	return EventEventTypeEnum{
		KAFKA: EventEventType{
			value: "KAFKA",
		},
		DIS: EventEventType{
			value: "DIS",
		},
		OBS: EventEventType{
			value: "OBS",
		},
	}
}

func (c EventEventType) Value() string {
	return c.value
}

func (c EventEventType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EventEventType) UnmarshalJSON(b []byte) error {
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

type EventFailPolicy struct {
	value string
}

type EventFailPolicyEnum struct {
	SUSPEND EventFailPolicy
	IGNORE  EventFailPolicy
}

func GetEventFailPolicyEnum() EventFailPolicyEnum {
	return EventFailPolicyEnum{
		SUSPEND: EventFailPolicy{
			value: "SUSPEND",
		},
		IGNORE: EventFailPolicy{
			value: "IGNORE",
		},
	}
}

func (c EventFailPolicy) Value() string {
	return c.value
}

func (c EventFailPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EventFailPolicy) UnmarshalJSON(b []byte) error {
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

type EventReadPolicy struct {
	value string
}

type EventReadPolicyEnum struct {
	LAST EventReadPolicy
	NEW  EventReadPolicy
}

func GetEventReadPolicyEnum() EventReadPolicyEnum {
	return EventReadPolicyEnum{
		LAST: EventReadPolicy{
			value: "LAST",
		},
		NEW: EventReadPolicy{
			value: "NEW",
		},
	}
}

func (c EventReadPolicy) Value() string {
	return c.value
}

func (c EventReadPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EventReadPolicy) UnmarshalJSON(b []byte) error {
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
