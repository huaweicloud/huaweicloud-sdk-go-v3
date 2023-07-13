package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowEventStreamingResponse Response Object
type ShowEventStreamingResponse struct {

	// 事件流名称，租户下唯一，由字母、数字、点、下划线和中划线组成，必须字母或数字开头
	Name string `json:"name"`

	// 事件流描述
	Description *string `json:"description,omitempty"`

	Source *EventStreamingSource `json:"source"`

	Sink *EventStreamingSink `json:"sink"`

	RuleConfig *EventStreamingCreateReqRuleConfig `json:"rule_config,omitempty"`

	Option *RunOption `json:"option,omitempty"`

	// 事件流状态
	Status *ShowEventStreamingResponseStatus `json:"status,omitempty"`

	// 事件流ID
	StreamingId *string `json:"streaming_id,omitempty"`

	// 创建时间
	CreatedTime *string `json:"created_time,omitempty"`

	// 更新时间
	UpdatedTime *string `json:"updated_time,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowEventStreamingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEventStreamingResponse struct{}"
	}

	return strings.Join([]string{"ShowEventStreamingResponse", string(data)}, " ")
}

type ShowEventStreamingResponseStatus struct {
	value string
}

type ShowEventStreamingResponseStatusEnum struct {
	CREATED ShowEventStreamingResponseStatus
	RUNNING ShowEventStreamingResponseStatus
	ERROR   ShowEventStreamingResponseStatus
	STOPPED ShowEventStreamingResponseStatus
}

func GetShowEventStreamingResponseStatusEnum() ShowEventStreamingResponseStatusEnum {
	return ShowEventStreamingResponseStatusEnum{
		CREATED: ShowEventStreamingResponseStatus{
			value: "CREATED",
		},
		RUNNING: ShowEventStreamingResponseStatus{
			value: "RUNNING",
		},
		ERROR: ShowEventStreamingResponseStatus{
			value: "ERROR",
		},
		STOPPED: ShowEventStreamingResponseStatus{
			value: "STOPPED",
		},
	}
}

func (c ShowEventStreamingResponseStatus) Value() string {
	return c.value
}

func (c ShowEventStreamingResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowEventStreamingResponseStatus) UnmarshalJSON(b []byte) error {
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
