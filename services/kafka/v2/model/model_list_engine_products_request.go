package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListEngineProductsRequest Request Object
type ListEngineProductsRequest struct {

	// **参数解释**： 消息引擎的类型。 **约束限制**： 不涉及。 **取值范围**： kafka **默认取值**： 不涉及。
	Engine ListEngineProductsRequestEngine `json:"engine"`

	// **参数解释**： 产品ID。 **约束限制**： 不涉及。 **取值范围**： [- s6.2u4g.single.small：对应规格kafka.2u4g.single.small。](tag:ax,dt,hws,hws_eu,hws_hk,srg) [- c6.2u4g.single：对应规格kafka.2u4g.single。](tag:ax,dt,hws,hws_eu,hws_hk,srg) [- s6.2u4g.cluster.small：对应规格kafka.2u4g.cluster.small。](tag:ax,dt,hws,hws_eu,hws_hk,srg) - c6.2u4g.cluster：对应规格kafka.2u4g.cluster。 - c6.4u8g.cluster：对应规格kafka.4u8g.cluster。 - c6.8u16g.cluster：对应规格kafka.8u16g.cluster。 - c6.12u24g.cluster：对应规格kafka.12u24g.cluster。 - c6.16u32g.cluster：对应规格kafka.16u32g.cluster。 **默认取值**： 不涉及。
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
	KAFKA ListEngineProductsRequestEngine
}

func GetListEngineProductsRequestEngineEnum() ListEngineProductsRequestEngineEnum {
	return ListEngineProductsRequestEngineEnum{
		KAFKA: ListEngineProductsRequestEngine{
			value: "kafka",
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
