package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ResizeEngineInstanceRequest Request Object
type ResizeEngineInstanceRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 消息引擎的类型。支持的类型为rabbitmq。
	Engine ResizeEngineInstanceRequestEngine `json:"engine"`

	Body *ResizeEngineInstanceReq `json:"body,omitempty"`
}

func (o ResizeEngineInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeEngineInstanceRequest struct{}"
	}

	return strings.Join([]string{"ResizeEngineInstanceRequest", string(data)}, " ")
}

type ResizeEngineInstanceRequestEngine struct {
	value string
}

type ResizeEngineInstanceRequestEngineEnum struct {
	RABBITMQ ResizeEngineInstanceRequestEngine
}

func GetResizeEngineInstanceRequestEngineEnum() ResizeEngineInstanceRequestEngineEnum {
	return ResizeEngineInstanceRequestEngineEnum{
		RABBITMQ: ResizeEngineInstanceRequestEngine{
			value: "rabbitmq",
		},
	}
}

func (c ResizeEngineInstanceRequestEngine) Value() string {
	return c.value
}

func (c ResizeEngineInstanceRequestEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResizeEngineInstanceRequestEngine) UnmarshalJSON(b []byte) error {
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
