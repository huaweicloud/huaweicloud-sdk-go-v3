package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CountResourcesByTagsRequest Request Object
type CountResourcesByTagsRequest struct {

	// 内容类型
	ContentType string `json:"Content-Type"`

	// 资源类型
	ResourceType CountResourcesByTagsRequestResourceType `json:"resource_type"`

	Body *QueryResourcesRequestBody `json:"body,omitempty"`
}

func (o CountResourcesByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountResourcesByTagsRequest struct{}"
	}

	return strings.Join([]string{"CountResourcesByTagsRequest", string(data)}, " ")
}

type CountResourcesByTagsRequestResourceType struct {
	value string
}

type CountResourcesByTagsRequestResourceTypeEnum struct {
	VPN_GATEWAY      CountResourcesByTagsRequestResourceType
	VPN_CONNECTION   CountResourcesByTagsRequestResourceType
	CUSTOMER_GATEWAY CountResourcesByTagsRequestResourceType
	P2C_VPN_GATEWAYS CountResourcesByTagsRequestResourceType
}

func GetCountResourcesByTagsRequestResourceTypeEnum() CountResourcesByTagsRequestResourceTypeEnum {
	return CountResourcesByTagsRequestResourceTypeEnum{
		VPN_GATEWAY: CountResourcesByTagsRequestResourceType{
			value: "vpn-gateway",
		},
		VPN_CONNECTION: CountResourcesByTagsRequestResourceType{
			value: "vpn-connection",
		},
		CUSTOMER_GATEWAY: CountResourcesByTagsRequestResourceType{
			value: "customer-gateway",
		},
		P2C_VPN_GATEWAYS: CountResourcesByTagsRequestResourceType{
			value: "p2c-vpn-gateways",
		},
	}
}

func (c CountResourcesByTagsRequestResourceType) Value() string {
	return c.value
}

func (c CountResourcesByTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CountResourcesByTagsRequestResourceType) UnmarshalJSON(b []byte) error {
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
