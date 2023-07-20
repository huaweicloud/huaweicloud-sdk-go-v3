package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateResourceTagRequest Request Object
type CreateResourceTagRequest struct {

	// 资源ID
	ResourceId string `json:"resource_id"`

	// 标签资源类型: - instance: 企业路由器实例 - route-table: 路由表 - vpc-attachment: VPC连接 - vgw-attachment: 虚拟网关连接 - peering-attachment: 对等连接（Peering）连接 - vpn-attachment: VPN网关连接 -  -  -  -  - attachments: 所有连接类型
	ResourceType CreateResourceTagRequestResourceType `json:"resource_type"`

	Body *CreateResourceTagRequestBody `json:"body,omitempty"`
}

func (o CreateResourceTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceTagRequest struct{}"
	}

	return strings.Join([]string{"CreateResourceTagRequest", string(data)}, " ")
}

type CreateResourceTagRequestResourceType struct {
	value string
}

type CreateResourceTagRequestResourceTypeEnum struct {
	INSTANCE           CreateResourceTagRequestResourceType
	ROUTE_TABLE        CreateResourceTagRequestResourceType
	VPC_ATTACHMENT     CreateResourceTagRequestResourceType
	DGW_ATTACHMENT     CreateResourceTagRequestResourceType
	VGW_ATTACHMENT     CreateResourceTagRequestResourceType
	PEERING_ATTACHMENT CreateResourceTagRequestResourceType
	VPN_ATTACHMENT     CreateResourceTagRequestResourceType
	CAN_ATTACHMENT     CreateResourceTagRequestResourceType
	ECN_ATTACHMENT     CreateResourceTagRequestResourceType
	GDGW_ATTACHMENT    CreateResourceTagRequestResourceType
	CONNECT_ATTACHMENT CreateResourceTagRequestResourceType
	CFW_ATTACHMENT     CreateResourceTagRequestResourceType
	ATTACHMENTS        CreateResourceTagRequestResourceType
}

func GetCreateResourceTagRequestResourceTypeEnum() CreateResourceTagRequestResourceTypeEnum {
	return CreateResourceTagRequestResourceTypeEnum{
		INSTANCE: CreateResourceTagRequestResourceType{
			value: "instance",
		},
		ROUTE_TABLE: CreateResourceTagRequestResourceType{
			value: "route-table",
		},
		VPC_ATTACHMENT: CreateResourceTagRequestResourceType{
			value: "vpc-attachment",
		},
		DGW_ATTACHMENT: CreateResourceTagRequestResourceType{
			value: "dgw-attachment",
		},
		VGW_ATTACHMENT: CreateResourceTagRequestResourceType{
			value: "vgw-attachment",
		},
		PEERING_ATTACHMENT: CreateResourceTagRequestResourceType{
			value: "peering-attachment",
		},
		VPN_ATTACHMENT: CreateResourceTagRequestResourceType{
			value: "vpn-attachment",
		},
		CAN_ATTACHMENT: CreateResourceTagRequestResourceType{
			value: "can-attachment",
		},
		ECN_ATTACHMENT: CreateResourceTagRequestResourceType{
			value: "ecn-attachment",
		},
		GDGW_ATTACHMENT: CreateResourceTagRequestResourceType{
			value: "gdgw-attachment",
		},
		CONNECT_ATTACHMENT: CreateResourceTagRequestResourceType{
			value: "connect-attachment",
		},
		CFW_ATTACHMENT: CreateResourceTagRequestResourceType{
			value: "cfw-attachment",
		},
		ATTACHMENTS: CreateResourceTagRequestResourceType{
			value: "attachments",
		},
	}
}

func (c CreateResourceTagRequestResourceType) Value() string {
	return c.value
}

func (c CreateResourceTagRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateResourceTagRequestResourceType) UnmarshalJSON(b []byte) error {
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
