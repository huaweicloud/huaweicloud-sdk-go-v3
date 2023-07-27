package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateLogStreamResponse Response Object
type UpdateLogStreamResponse struct {

	// 创建该日志流的时间
	CreationTime *int64 `json:"creation_time,omitempty"`

	// 日志流的名称。
	LogTopicName *string `json:"log_topic_name,omitempty"`

	// 日志流ID。
	LogTopicId *string `json:"log_topic_id,omitempty"`

	// 日志存储时间（天）。
	TtlInDays      *UpdateLogStreamResponseTtlInDays `json:"ttl_in_days,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o UpdateLogStreamResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLogStreamResponse struct{}"
	}

	return strings.Join([]string{"UpdateLogStreamResponse", string(data)}, " ")
}

type UpdateLogStreamResponseTtlInDays struct {
	value int32
}

type UpdateLogStreamResponseTtlInDaysEnum struct {
	E_7 UpdateLogStreamResponseTtlInDays
}

func GetUpdateLogStreamResponseTtlInDaysEnum() UpdateLogStreamResponseTtlInDaysEnum {
	return UpdateLogStreamResponseTtlInDaysEnum{
		E_7: UpdateLogStreamResponseTtlInDays{
			value: 7,
		},
	}
}

func (c UpdateLogStreamResponseTtlInDays) Value() int32 {
	return c.value
}

func (c UpdateLogStreamResponseTtlInDays) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateLogStreamResponseTtlInDays) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
