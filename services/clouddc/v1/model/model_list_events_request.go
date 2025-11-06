package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListEventsRequest Request Object
type ListEventsRequest struct {

	// 分页游标
	Offset *int32 `json:"offset,omitempty"`

	// 分页大小
	Limit *int32 `json:"limit,omitempty"`

	// 事件级别，值为critical为紧急，major为重要，minor为次要，info为提示
	EventLevel *ListEventsRequestEventLevel `json:"event_level,omitempty"`

	// 告警资源ID
	ResourceId *string `json:"resource_id,omitempty"`

	// 查询开始时间，格式为时间戳（毫秒），from 必须小于 to，且查询时间范围最大不超过30天。
	From string `json:"from"`

	// 查询截止时间，格式为时间戳（毫秒），from 必须小于 to，且查询时间范围最大不超过30天。
	To string `json:"to"`
}

func (o ListEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventsRequest struct{}"
	}

	return strings.Join([]string{"ListEventsRequest", string(data)}, " ")
}

type ListEventsRequestEventLevel struct {
	value string
}

type ListEventsRequestEventLevelEnum struct {
	CRITICAL ListEventsRequestEventLevel
	MAJOR    ListEventsRequestEventLevel
	MINOR    ListEventsRequestEventLevel
	INFO     ListEventsRequestEventLevel
}

func GetListEventsRequestEventLevelEnum() ListEventsRequestEventLevelEnum {
	return ListEventsRequestEventLevelEnum{
		CRITICAL: ListEventsRequestEventLevel{
			value: "critical",
		},
		MAJOR: ListEventsRequestEventLevel{
			value: "major",
		},
		MINOR: ListEventsRequestEventLevel{
			value: "minor",
		},
		INFO: ListEventsRequestEventLevel{
			value: "info",
		},
	}
}

func (c ListEventsRequestEventLevel) Value() string {
	return c.value
}

func (c ListEventsRequestEventLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEventsRequestEventLevel) UnmarshalJSON(b []byte) error {
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
