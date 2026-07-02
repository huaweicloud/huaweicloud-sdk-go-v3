package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowKafkaLogTaskRequest Request Object
type ShowKafkaLogTaskRequest struct {

	// **参数解释**： 实例ID。获取方法如下：调用“查询所有实例列表”接口，从响应体中获取实例ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**： 日志任务类型。 **约束限制**： 不涉及。 **取值范围**： - REBALANCE：重平衡日志。 - topic_log：Topic日志。 **默认取值**： 不涉及。
	LogType ShowKafkaLogTaskRequestLogType `json:"log_type"`
}

func (o ShowKafkaLogTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKafkaLogTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowKafkaLogTaskRequest", string(data)}, " ")
}

type ShowKafkaLogTaskRequestLogType struct {
	value string
}

type ShowKafkaLogTaskRequestLogTypeEnum struct {
	REBALANCE ShowKafkaLogTaskRequestLogType
	TOPIC_LOG ShowKafkaLogTaskRequestLogType
}

func GetShowKafkaLogTaskRequestLogTypeEnum() ShowKafkaLogTaskRequestLogTypeEnum {
	return ShowKafkaLogTaskRequestLogTypeEnum{
		REBALANCE: ShowKafkaLogTaskRequestLogType{
			value: "REBALANCE",
		},
		TOPIC_LOG: ShowKafkaLogTaskRequestLogType{
			value: "topic_log",
		},
	}
}

func (c ShowKafkaLogTaskRequestLogType) Value() string {
	return c.value
}

func (c ShowKafkaLogTaskRequestLogType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowKafkaLogTaskRequestLogType) UnmarshalJSON(b []byte) error {
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
