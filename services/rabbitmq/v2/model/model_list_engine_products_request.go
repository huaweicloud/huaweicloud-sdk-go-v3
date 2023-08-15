package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListEngineProductsRequest Request Object
type ListEngineProductsRequest struct {

	// 消息引擎的类型。
	Engine ListEngineProductsRequestEngine `json:"engine"`

	// 产品ID。
	ProductId *string `json:"product_id,omitempty"`
}

func (o ListEngineProductsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEngineProductsRequest struct{}"
	}

	return strings.Join([]string{"ListEngineProductsRequest", string(data)}, " ")
}

type ListEngineProductsRequestEngine struct {
	value string
}

type ListEngineProductsRequestEngineEnum struct {
	RABBITMQ ListEngineProductsRequestEngine
}

func GetListEngineProductsRequestEngineEnum() ListEngineProductsRequestEngineEnum {
	return ListEngineProductsRequestEngineEnum{
		RABBITMQ: ListEngineProductsRequestEngine{
			value: "rabbitmq",
		},
	}
}

func (c ListEngineProductsRequestEngine) Value() string {
	return c.value
}

func (c ListEngineProductsRequestEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEngineProductsRequestEngine) UnmarshalJSON(b []byte) error {
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
