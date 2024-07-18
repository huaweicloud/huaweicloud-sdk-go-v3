package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchDeleteResourceTagsRequest Request Object
type BatchDeleteResourceTagsRequest struct {

	// 内容类型
	ContentType string `json:"Content-Type"`

	// 资源类型
	ResourceType BatchDeleteResourceTagsRequestResourceType `json:"resource_type"`

	// 资源ID
	ResourceId string `json:"resource_id"`

	Body *DeleteResourcesTagsRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteResourceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteResourceTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteResourceTagsRequest", string(data)}, " ")
}

type BatchDeleteResourceTagsRequestResourceType struct {
	value string
}

type BatchDeleteResourceTagsRequestResourceTypeEnum struct {
	VPN_GATEWAY      BatchDeleteResourceTagsRequestResourceType
	VPN_CONNECTION   BatchDeleteResourceTagsRequestResourceType
	CUSTOMER_GATEWAY BatchDeleteResourceTagsRequestResourceType
	P2C_VPN_GATEWAYS BatchDeleteResourceTagsRequestResourceType
}

func GetBatchDeleteResourceTagsRequestResourceTypeEnum() BatchDeleteResourceTagsRequestResourceTypeEnum {
	return BatchDeleteResourceTagsRequestResourceTypeEnum{
		VPN_GATEWAY: BatchDeleteResourceTagsRequestResourceType{
			value: "vpn-gateway",
		},
		VPN_CONNECTION: BatchDeleteResourceTagsRequestResourceType{
			value: "vpn-connection",
		},
		CUSTOMER_GATEWAY: BatchDeleteResourceTagsRequestResourceType{
			value: "customer-gateway",
		},
		P2C_VPN_GATEWAYS: BatchDeleteResourceTagsRequestResourceType{
			value: "p2c-vpn-gateways",
		},
	}
}

func (c BatchDeleteResourceTagsRequestResourceType) Value() string {
	return c.value
}

func (c BatchDeleteResourceTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteResourceTagsRequestResourceType) UnmarshalJSON(b []byte) error {
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
