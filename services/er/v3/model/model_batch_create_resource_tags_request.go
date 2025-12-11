package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchCreateResourceTagsRequest Request Object
type BatchCreateResourceTagsRequest struct {

	// 资源ID
	ResourceId string `json:"resource_id"`

	// 标签资源类型: - instance: 企业路由器实例 - route-table: 路由表 - vpc-attachment: VPC连接 - vgw-attachment: 虚拟网关连接 - peering-attachment: 对等连接（Peering）连接 - vpn-attachment: VPN网关连接 - attachments: 所有连接类型
	ResourceType BatchCreateResourceTagsRequestResourceType `json:"resource_type"`

	Body *BatchOperateResourceTagsRequestBody `json:"body,omitempty"`
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
	INSTANCE           BatchCreateResourceTagsRequestResourceType
	ROUTE_TABLE        BatchCreateResourceTagsRequestResourceType
	VPC_ATTACHMENT     BatchCreateResourceTagsRequestResourceType
	DGW_ATTACHMENT     BatchCreateResourceTagsRequestResourceType
	VGW_ATTACHMENT     BatchCreateResourceTagsRequestResourceType
	PEERING_ATTACHMENT BatchCreateResourceTagsRequestResourceType
	VPN_ATTACHMENT     BatchCreateResourceTagsRequestResourceType
	CAN_ATTACHMENT     BatchCreateResourceTagsRequestResourceType
	ECN_ATTACHMENT     BatchCreateResourceTagsRequestResourceType
	GDGW_ATTACHMENT    BatchCreateResourceTagsRequestResourceType
	CONNECT_ATTACHMENT BatchCreateResourceTagsRequestResourceType
	CFW_ATTACHMENT     BatchCreateResourceTagsRequestResourceType
	ATTACHMENTS        BatchCreateResourceTagsRequestResourceType
}

func GetBatchCreateResourceTagsRequestResourceTypeEnum() BatchCreateResourceTagsRequestResourceTypeEnum {
	return BatchCreateResourceTagsRequestResourceTypeEnum{
		INSTANCE: BatchCreateResourceTagsRequestResourceType{
			value: "instance",
		},
		ROUTE_TABLE: BatchCreateResourceTagsRequestResourceType{
			value: "route-table",
		},
		VPC_ATTACHMENT: BatchCreateResourceTagsRequestResourceType{
			value: "vpc-attachment",
		},
		DGW_ATTACHMENT: BatchCreateResourceTagsRequestResourceType{
			value: "dgw-attachment",
		},
		VGW_ATTACHMENT: BatchCreateResourceTagsRequestResourceType{
			value: "vgw-attachment",
		},
		PEERING_ATTACHMENT: BatchCreateResourceTagsRequestResourceType{
			value: "peering-attachment",
		},
		VPN_ATTACHMENT: BatchCreateResourceTagsRequestResourceType{
			value: "vpn-attachment",
		},
		CAN_ATTACHMENT: BatchCreateResourceTagsRequestResourceType{
			value: "can-attachment",
		},
		ECN_ATTACHMENT: BatchCreateResourceTagsRequestResourceType{
			value: "ecn-attachment",
		},
		GDGW_ATTACHMENT: BatchCreateResourceTagsRequestResourceType{
			value: "gdgw-attachment",
		},
		CONNECT_ATTACHMENT: BatchCreateResourceTagsRequestResourceType{
			value: "connect-attachment",
		},
		CFW_ATTACHMENT: BatchCreateResourceTagsRequestResourceType{
			value: "cfw-attachment",
		},
		ATTACHMENTS: BatchCreateResourceTagsRequestResourceType{
			value: "attachments",
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
