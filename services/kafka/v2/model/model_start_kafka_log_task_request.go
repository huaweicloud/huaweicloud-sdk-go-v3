package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StartKafkaLogTaskRequest Request Object
type StartKafkaLogTaskRequest struct {

	// **参数解释**： 实例ID。获取方法如下：调用“查询所有实例列表”接口，从响应体中获取实例ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**： 日志任务类型。 **约束限制**： 不涉及。 **取值范围**： - REBALANCE：重平衡日志。 - topic_log：Topic日志。 **默认取值**： 不涉及。
	LogType StartKafkaLogTaskRequestLogType `json:"log_type"`

	Body *StartKafkaLogTaskReq `json:"body,omitempty"`
}

func (o StartKafkaLogTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartKafkaLogTaskRequest struct{}"
	}

	return strings.Join([]string{"StartKafkaLogTaskRequest", string(data)}, " ")
}

type StartKafkaLogTaskRequestLogType struct {
	value string
}

type StartKafkaLogTaskRequestLogTypeEnum struct {
	REBALANCE StartKafkaLogTaskRequestLogType
	TOPIC_LOG StartKafkaLogTaskRequestLogType
}

func GetStartKafkaLogTaskRequestLogTypeEnum() StartKafkaLogTaskRequestLogTypeEnum {
	return StartKafkaLogTaskRequestLogTypeEnum{
		REBALANCE: StartKafkaLogTaskRequestLogType{
			value: "REBALANCE",
		},
		TOPIC_LOG: StartKafkaLogTaskRequestLogType{
			value: "topic_log",
		},
	}
}

func (c StartKafkaLogTaskRequestLogType) Value() string {
	return c.value
}

func (c StartKafkaLogTaskRequestLogType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StartKafkaLogTaskRequestLogType) UnmarshalJSON(b []byte) error {
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
