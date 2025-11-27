package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListNatGatewaySpecsResponse Response Object
type ListNatGatewaySpecsResponse struct {

	// 可创建的NAT网关实例列表  取值范围：  “1”：小型，SNAT最大连接数10000  “2”：中型，SNAT最大连接数50000  “3”：大型，SNAT最大连接数200000  “4”：超大型，SNAT最大连接数1000000  “5”：企业型，SNAT最大连接数10000000
	Specs          *[]ListNatGatewaySpecsResponseSpecs `json:"specs,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o ListNatGatewaySpecsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNatGatewaySpecsResponse struct{}"
	}

	return strings.Join([]string{"ListNatGatewaySpecsResponse", string(data)}, " ")
}

type ListNatGatewaySpecsResponseSpecs struct {
	value string
}

type ListNatGatewaySpecsResponseSpecsEnum struct {
	E_1 ListNatGatewaySpecsResponseSpecs
	E_2 ListNatGatewaySpecsResponseSpecs
	E_3 ListNatGatewaySpecsResponseSpecs
	E_4 ListNatGatewaySpecsResponseSpecs
	E_5 ListNatGatewaySpecsResponseSpecs
}

func GetListNatGatewaySpecsResponseSpecsEnum() ListNatGatewaySpecsResponseSpecsEnum {
	return ListNatGatewaySpecsResponseSpecsEnum{
		E_1: ListNatGatewaySpecsResponseSpecs{
			value: "1",
		},
		E_2: ListNatGatewaySpecsResponseSpecs{
			value: "2",
		},
		E_3: ListNatGatewaySpecsResponseSpecs{
			value: "3",
		},
		E_4: ListNatGatewaySpecsResponseSpecs{
			value: "4",
		},
		E_5: ListNatGatewaySpecsResponseSpecs{
			value: "5",
		},
	}
}

func (c ListNatGatewaySpecsResponseSpecs) Value() string {
	return c.value
}

func (c ListNatGatewaySpecsResponseSpecs) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListNatGatewaySpecsResponseSpecs) UnmarshalJSON(b []byte) error {
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
