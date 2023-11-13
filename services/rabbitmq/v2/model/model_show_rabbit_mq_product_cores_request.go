package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowRabbitMqProductCoresRequest Request Object
type ShowRabbitMqProductCoresRequest struct {

	// 消息引擎的类型。
	Engine ShowRabbitMqProductCoresRequestEngine `json:"engine"`

	// 产品ID。
	ProductId string `json:"product_id"`

	// 代理个数。  当产品为单机类型，代理个数只能为1；当产品为集群类型，可选3、5、7个代理个数。  产品类型为single时:   - 1  产品类型为cluster时:   - 3   - 5   - 7
	BrokerNum int32 `json:"broker_num"`

	// 实例ID。
	InstanceId *string `json:"instance_id,omitempty"`
}

func (o ShowRabbitMqProductCoresRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRabbitMqProductCoresRequest struct{}"
	}

	return strings.Join([]string{"ShowRabbitMqProductCoresRequest", string(data)}, " ")
}

type ShowRabbitMqProductCoresRequestEngine struct {
	value string
}

type ShowRabbitMqProductCoresRequestEngineEnum struct {
	RABBITMQ ShowRabbitMqProductCoresRequestEngine
}

func GetShowRabbitMqProductCoresRequestEngineEnum() ShowRabbitMqProductCoresRequestEngineEnum {
	return ShowRabbitMqProductCoresRequestEngineEnum{
		RABBITMQ: ShowRabbitMqProductCoresRequestEngine{
			value: "rabbitmq",
		},
	}
}

func (c ShowRabbitMqProductCoresRequestEngine) Value() string {
	return c.value
}

func (c ShowRabbitMqProductCoresRequestEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowRabbitMqProductCoresRequestEngine) UnmarshalJSON(b []byte) error {
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
