package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateKafkaPortProtocolRequest Request Object
type UpdateKafkaPortProtocolRequest struct {

	// 消息引擎。
	Engine UpdateKafkaPortProtocolRequestEngine `json:"engine"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *PlainSslEnableRequest `json:"body,omitempty"`
}

func (o UpdateKafkaPortProtocolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKafkaPortProtocolRequest struct{}"
	}

	return strings.Join([]string{"UpdateKafkaPortProtocolRequest", string(data)}, " ")
}

type UpdateKafkaPortProtocolRequestEngine struct {
	value string
}

type UpdateKafkaPortProtocolRequestEngineEnum struct {
	KAFKA UpdateKafkaPortProtocolRequestEngine
}

func GetUpdateKafkaPortProtocolRequestEngineEnum() UpdateKafkaPortProtocolRequestEngineEnum {
	return UpdateKafkaPortProtocolRequestEngineEnum{
		KAFKA: UpdateKafkaPortProtocolRequestEngine{
			value: "kafka",
		},
	}
}

func (c UpdateKafkaPortProtocolRequestEngine) Value() string {
	return c.value
}

func (c UpdateKafkaPortProtocolRequestEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateKafkaPortProtocolRequestEngine) UnmarshalJSON(b []byte) error {
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
