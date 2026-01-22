package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateInstanceByEngineRequest Request Object
type CreateInstanceByEngineRequest struct {

	// **参数解释**： 消息引擎。 **约束限制**： 不涉及。 **取值范围**： - rocketmq：RocketMQ消息引擎。 - reliability：RocketMQ消息引擎别称。 **默认取值**： 不涉及。
	Engine CreateInstanceByEngineRequestEngine `json:"engine"`

	Body *CreateInstanceByEngineReq `json:"body,omitempty"`
}

func (o CreateInstanceByEngineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceByEngineRequest struct{}"
	}

	return strings.Join([]string{"CreateInstanceByEngineRequest", string(data)}, " ")
}

type CreateInstanceByEngineRequestEngine struct {
	value string
}

type CreateInstanceByEngineRequestEngineEnum struct {
	ROCKETMQ    CreateInstanceByEngineRequestEngine
	RELIABILITY CreateInstanceByEngineRequestEngine
}

func GetCreateInstanceByEngineRequestEngineEnum() CreateInstanceByEngineRequestEngineEnum {
	return CreateInstanceByEngineRequestEngineEnum{
		ROCKETMQ: CreateInstanceByEngineRequestEngine{
			value: "rocketmq",
		},
		RELIABILITY: CreateInstanceByEngineRequestEngine{
			value: "reliability",
		},
	}
}

func (c CreateInstanceByEngineRequestEngine) Value() string {
	return c.value
}

func (c CreateInstanceByEngineRequestEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateInstanceByEngineRequestEngine) UnmarshalJSON(b []byte) error {
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
