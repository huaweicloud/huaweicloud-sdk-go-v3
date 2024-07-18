package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListProjectTagsRequest Request Object
type ListProjectTagsRequest struct {

	// 内容类型
	ContentType string `json:"Content-Type"`

	// 资源类型
	ResourceType ListProjectTagsRequestResourceType `json:"resource_type"`
}

func (o ListProjectTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectTagsRequest struct{}"
	}

	return strings.Join([]string{"ListProjectTagsRequest", string(data)}, " ")
}

type ListProjectTagsRequestResourceType struct {
	value string
}

type ListProjectTagsRequestResourceTypeEnum struct {
	VPN_GATEWAY      ListProjectTagsRequestResourceType
	VPN_CONNECTION   ListProjectTagsRequestResourceType
	CUSTOMER_GATEWAY ListProjectTagsRequestResourceType
	P2C_VPN_GATEWAYS ListProjectTagsRequestResourceType
}

func GetListProjectTagsRequestResourceTypeEnum() ListProjectTagsRequestResourceTypeEnum {
	return ListProjectTagsRequestResourceTypeEnum{
		VPN_GATEWAY: ListProjectTagsRequestResourceType{
			value: "vpn-gateway",
		},
		VPN_CONNECTION: ListProjectTagsRequestResourceType{
			value: "vpn-connection",
		},
		CUSTOMER_GATEWAY: ListProjectTagsRequestResourceType{
			value: "customer-gateway",
		},
		P2C_VPN_GATEWAYS: ListProjectTagsRequestResourceType{
			value: "p2c-vpn-gateways",
		},
	}
}

func (c ListProjectTagsRequestResourceType) Value() string {
	return c.value
}

func (c ListProjectTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListProjectTagsRequestResourceType) UnmarshalJSON(b []byte) error {
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
