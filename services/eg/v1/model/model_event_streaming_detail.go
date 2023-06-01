package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type EventStreamingDetail struct {

	// 事件流名称，租户下唯一，由字母、数字、点、下划线和中划线组成，必须字母或数字开头
	Name string `json:"name"`

	// 事件流描述
	Description *string `json:"description,omitempty"`

	Source *EventStreamingSource `json:"source"`

	Sink *EventStreamingSink `json:"sink"`

	RuleConfig *EventStreamingCreateReqRuleConfig `json:"rule_config,omitempty"`

	Option *RunOption `json:"option,omitempty"`

	// 事件流状态
	Status *EventStreamingDetailStatus `json:"status,omitempty"`

	// 事件流ID
	StreamingId *string `json:"streaming_id,omitempty"`

	// 创建时间
	CreatedTime *string `json:"created_time,omitempty"`

	// 更新时间
	UpdatedTime *string `json:"updated_time,omitempty"`
}

func (o EventStreamingDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventStreamingDetail struct{}"
	}

	return strings.Join([]string{"EventStreamingDetail", string(data)}, " ")
}

type EventStreamingDetailStatus struct {
	value string
}

type EventStreamingDetailStatusEnum struct {
	CREATED EventStreamingDetailStatus
	RUNNING EventStreamingDetailStatus
	ERROR   EventStreamingDetailStatus
	STOPPED EventStreamingDetailStatus
}

func GetEventStreamingDetailStatusEnum() EventStreamingDetailStatusEnum {
	return EventStreamingDetailStatusEnum{
		CREATED: EventStreamingDetailStatus{
			value: "CREATED",
		},
		RUNNING: EventStreamingDetailStatus{
			value: "RUNNING",
		},
		ERROR: EventStreamingDetailStatus{
			value: "ERROR",
		},
		STOPPED: EventStreamingDetailStatus{
			value: "STOPPED",
		},
	}
}

func (c EventStreamingDetailStatus) Value() string {
	return c.value
}

func (c EventStreamingDetailStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EventStreamingDetailStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
