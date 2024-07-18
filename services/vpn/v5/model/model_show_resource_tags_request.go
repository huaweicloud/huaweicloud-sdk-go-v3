package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowResourceTagsRequest Request Object
type ShowResourceTagsRequest struct {

	// 内容类型
	ContentType string `json:"Content-Type"`

	// 资源类型
	ResourceType ShowResourceTagsRequestResourceType `json:"resource_type"`

	// 资源ID
	ResourceId string `json:"resource_id"`
}

func (o ShowResourceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowResourceTagsRequest", string(data)}, " ")
}

type ShowResourceTagsRequestResourceType struct {
	value string
}

type ShowResourceTagsRequestResourceTypeEnum struct {
	VPN_GATEWAY      ShowResourceTagsRequestResourceType
	VPN_CONNECTION   ShowResourceTagsRequestResourceType
	CUSTOMER_GATEWAY ShowResourceTagsRequestResourceType
	P2C_VPN_GATEWAYS ShowResourceTagsRequestResourceType
}

func GetShowResourceTagsRequestResourceTypeEnum() ShowResourceTagsRequestResourceTypeEnum {
	return ShowResourceTagsRequestResourceTypeEnum{
		VPN_GATEWAY: ShowResourceTagsRequestResourceType{
			value: "vpn-gateway",
		},
		VPN_CONNECTION: ShowResourceTagsRequestResourceType{
			value: "vpn-connection",
		},
		CUSTOMER_GATEWAY: ShowResourceTagsRequestResourceType{
			value: "customer-gateway",
		},
		P2C_VPN_GATEWAYS: ShowResourceTagsRequestResourceType{
			value: "p2c-vpn-gateways",
		},
	}
}

func (c ShowResourceTagsRequestResourceType) Value() string {
	return c.value
}

func (c ShowResourceTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowResourceTagsRequestResourceType) UnmarshalJSON(b []byte) error {
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
