package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListResourcesByTagsRequest Request Object
type ListResourcesByTagsRequest struct {

	// 内容类型
	ContentType string `json:"Content-Type"`

	// 资源类型
	ResourceType ListResourcesByTagsRequestResourceType `json:"resource_type"`

	// limit
	Limit *string `json:"limit,omitempty"`

	// offset
	Offset *string `json:"offset,omitempty"`

	Body *QueryResourcesRequestBody `json:"body,omitempty"`
}

func (o ListResourcesByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourcesByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListResourcesByTagsRequest", string(data)}, " ")
}

type ListResourcesByTagsRequestResourceType struct {
	value string
}

type ListResourcesByTagsRequestResourceTypeEnum struct {
	VPN_GATEWAY      ListResourcesByTagsRequestResourceType
	VPN_CONNECTION   ListResourcesByTagsRequestResourceType
	CUSTOMER_GATEWAY ListResourcesByTagsRequestResourceType
	P2C_VPN_GATEWAYS ListResourcesByTagsRequestResourceType
}

func GetListResourcesByTagsRequestResourceTypeEnum() ListResourcesByTagsRequestResourceTypeEnum {
	return ListResourcesByTagsRequestResourceTypeEnum{
		VPN_GATEWAY: ListResourcesByTagsRequestResourceType{
			value: "vpn-gateway",
		},
		VPN_CONNECTION: ListResourcesByTagsRequestResourceType{
			value: "vpn-connection",
		},
		CUSTOMER_GATEWAY: ListResourcesByTagsRequestResourceType{
			value: "customer-gateway",
		},
		P2C_VPN_GATEWAYS: ListResourcesByTagsRequestResourceType{
			value: "p2c-vpn-gateways",
		},
	}
}

func (c ListResourcesByTagsRequestResourceType) Value() string {
	return c.value
}

func (c ListResourcesByTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListResourcesByTagsRequestResourceType) UnmarshalJSON(b []byte) error {
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
