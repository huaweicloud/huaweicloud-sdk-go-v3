package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ShowResourceTagRequest struct {

	// 资源ID
	ResourceId string `json:"resource_id"`

	// - instance: 企业路由器实例 - route-table: 路由表 - vpc-attachment: VPC连接 - vgw-attachment: 虚拟网关连接 - peering-attachment: 对等连接（Peering）连接 - vpn-attachment: VPN网关连接 - attachments: 所有连接类型
	ResourceType ShowResourceTagRequestResourceType `json:"resource_type"`
}

func (o ShowResourceTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceTagRequest struct{}"
	}

	return strings.Join([]string{"ShowResourceTagRequest", string(data)}, " ")
}

type ShowResourceTagRequestResourceType struct {
	value string
}

type ShowResourceTagRequestResourceTypeEnum struct {
	INSTANCE           ShowResourceTagRequestResourceType
	ROUTE_TABLE        ShowResourceTagRequestResourceType
	VPC_ATTACHMENT     ShowResourceTagRequestResourceType
	DGW_ATTACHMENT     ShowResourceTagRequestResourceType
	VGW_ATTACHMENT     ShowResourceTagRequestResourceType
	PEERING_ATTACHMENT ShowResourceTagRequestResourceType
	VPN_ATTACHMENT     ShowResourceTagRequestResourceType
	CAN_ATTACHMENT     ShowResourceTagRequestResourceType
	GDGW_ATTACHMENT    ShowResourceTagRequestResourceType
	ATTACHMENTS        ShowResourceTagRequestResourceType
}

func GetShowResourceTagRequestResourceTypeEnum() ShowResourceTagRequestResourceTypeEnum {
	return ShowResourceTagRequestResourceTypeEnum{
		INSTANCE: ShowResourceTagRequestResourceType{
			value: "instance",
		},
		ROUTE_TABLE: ShowResourceTagRequestResourceType{
			value: "route-table",
		},
		VPC_ATTACHMENT: ShowResourceTagRequestResourceType{
			value: "vpc-attachment",
		},
		DGW_ATTACHMENT: ShowResourceTagRequestResourceType{
			value: "dgw-attachment",
		},
		VGW_ATTACHMENT: ShowResourceTagRequestResourceType{
			value: "vgw-attachment",
		},
		PEERING_ATTACHMENT: ShowResourceTagRequestResourceType{
			value: "peering-attachment",
		},
		VPN_ATTACHMENT: ShowResourceTagRequestResourceType{
			value: "vpn-attachment",
		},
		CAN_ATTACHMENT: ShowResourceTagRequestResourceType{
			value: "can-attachment",
		},
		GDGW_ATTACHMENT: ShowResourceTagRequestResourceType{
			value: "gdgw-attachment",
		},
		ATTACHMENTS: ShowResourceTagRequestResourceType{
			value: "attachments",
		},
	}
}

func (c ShowResourceTagRequestResourceType) Value() string {
	return c.value
}

func (c ShowResourceTagRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowResourceTagRequestResourceType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
