package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteResourceTagRequest Request Object
type DeleteResourceTagRequest struct {

	// 标签键
	Key string `json:"key"`

	// 资源ID
	ResourceId string `json:"resource_id"`

	// 标签资源类型: - instance: 企业路由器实例 - route-table: 路由表 - vpc-attachment: VPC连接 - vgw-attachment: 虚拟网关连接 - peering-attachment: 对等连接（Peering）连接 - vpn-attachment: VPN网关连接 -  -  -  -  - attachments: 所有连接类型
	ResourceType DeleteResourceTagRequestResourceType `json:"resource_type"`
}

func (o DeleteResourceTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResourceTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteResourceTagRequest", string(data)}, " ")
}

type DeleteResourceTagRequestResourceType struct {
	value string
}

type DeleteResourceTagRequestResourceTypeEnum struct {
	INSTANCE           DeleteResourceTagRequestResourceType
	ROUTE_TABLE        DeleteResourceTagRequestResourceType
	VPC_ATTACHMENT     DeleteResourceTagRequestResourceType
	DGW_ATTACHMENT     DeleteResourceTagRequestResourceType
	VGW_ATTACHMENT     DeleteResourceTagRequestResourceType
	PEERING_ATTACHMENT DeleteResourceTagRequestResourceType
	VPN_ATTACHMENT     DeleteResourceTagRequestResourceType
	CAN_ATTACHMENT     DeleteResourceTagRequestResourceType
	ECN_ATTACHMENT     DeleteResourceTagRequestResourceType
	GDGW_ATTACHMENT    DeleteResourceTagRequestResourceType
	CONNECT_ATTACHMENT DeleteResourceTagRequestResourceType
	CFW_ATTACHMENT     DeleteResourceTagRequestResourceType
	ATTACHMENTS        DeleteResourceTagRequestResourceType
}

func GetDeleteResourceTagRequestResourceTypeEnum() DeleteResourceTagRequestResourceTypeEnum {
	return DeleteResourceTagRequestResourceTypeEnum{
		INSTANCE: DeleteResourceTagRequestResourceType{
			value: "instance",
		},
		ROUTE_TABLE: DeleteResourceTagRequestResourceType{
			value: "route-table",
		},
		VPC_ATTACHMENT: DeleteResourceTagRequestResourceType{
			value: "vpc-attachment",
		},
		DGW_ATTACHMENT: DeleteResourceTagRequestResourceType{
			value: "dgw-attachment",
		},
		VGW_ATTACHMENT: DeleteResourceTagRequestResourceType{
			value: "vgw-attachment",
		},
		PEERING_ATTACHMENT: DeleteResourceTagRequestResourceType{
			value: "peering-attachment",
		},
		VPN_ATTACHMENT: DeleteResourceTagRequestResourceType{
			value: "vpn-attachment",
		},
		CAN_ATTACHMENT: DeleteResourceTagRequestResourceType{
			value: "can-attachment",
		},
		ECN_ATTACHMENT: DeleteResourceTagRequestResourceType{
			value: "ecn-attachment",
		},
		GDGW_ATTACHMENT: DeleteResourceTagRequestResourceType{
			value: "gdgw-attachment",
		},
		CONNECT_ATTACHMENT: DeleteResourceTagRequestResourceType{
			value: "connect-attachment",
		},
		CFW_ATTACHMENT: DeleteResourceTagRequestResourceType{
			value: "cfw-attachment",
		},
		ATTACHMENTS: DeleteResourceTagRequestResourceType{
			value: "attachments",
		},
	}
}

func (c DeleteResourceTagRequestResourceType) Value() string {
	return c.value
}

func (c DeleteResourceTagRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteResourceTagRequestResourceType) UnmarshalJSON(b []byte) error {
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
