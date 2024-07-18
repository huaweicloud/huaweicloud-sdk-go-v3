package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchCreateResourceTagsRequest Request Object
type BatchCreateResourceTagsRequest struct {

	// 内容类型
	ContentType string `json:"Content-Type"`

	// 资源类型
	ResourceType BatchCreateResourceTagsRequestResourceType `json:"resource_type"`

	// 资源ID
	ResourceId string `json:"resource_id"`

	Body *CreateResourcesTagsRequestBody `json:"body,omitempty"`
}

func (o BatchCreateResourceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateResourceTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateResourceTagsRequest", string(data)}, " ")
}

type BatchCreateResourceTagsRequestResourceType struct {
	value string
}

type BatchCreateResourceTagsRequestResourceTypeEnum struct {
	VPN_GATEWAY      BatchCreateResourceTagsRequestResourceType
	VPN_CONNECTION   BatchCreateResourceTagsRequestResourceType
	CUSTOMER_GATEWAY BatchCreateResourceTagsRequestResourceType
	P2C_VPN_GATEWAYS BatchCreateResourceTagsRequestResourceType
}

func GetBatchCreateResourceTagsRequestResourceTypeEnum() BatchCreateResourceTagsRequestResourceTypeEnum {
	return BatchCreateResourceTagsRequestResourceTypeEnum{
		VPN_GATEWAY: BatchCreateResourceTagsRequestResourceType{
			value: "vpn-gateway",
		},
		VPN_CONNECTION: BatchCreateResourceTagsRequestResourceType{
			value: "vpn-connection",
		},
		CUSTOMER_GATEWAY: BatchCreateResourceTagsRequestResourceType{
			value: "customer-gateway",
		},
		P2C_VPN_GATEWAYS: BatchCreateResourceTagsRequestResourceType{
			value: "p2c-vpn-gateways",
		},
	}
}

func (c BatchCreateResourceTagsRequestResourceType) Value() string {
	return c.value
}

func (c BatchCreateResourceTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateResourceTagsRequestResourceType) UnmarshalJSON(b []byte) error {
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
