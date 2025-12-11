package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListProjectTagsRequest Request Object
type ListProjectTagsRequest struct {

	// 标签资源类型: - instance: 企业路由器实例 - route-table: 路由表 - vpc-attachment: VPC连接 - vgw-attachment: 虚拟网关连接 - peering-attachment: 对等连接（Peering）连接 - vpn-attachment: VPN网关连接 - attachments: 所有连接类型
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
	INSTANCE           ListProjectTagsRequestResourceType
	ROUTE_TABLE        ListProjectTagsRequestResourceType
	VPC_ATTACHMENT     ListProjectTagsRequestResourceType
	DGW_ATTACHMENT     ListProjectTagsRequestResourceType
	VGW_ATTACHMENT     ListProjectTagsRequestResourceType
	PEERING_ATTACHMENT ListProjectTagsRequestResourceType
	VPN_ATTACHMENT     ListProjectTagsRequestResourceType
	CAN_ATTACHMENT     ListProjectTagsRequestResourceType
	ECN_ATTACHMENT     ListProjectTagsRequestResourceType
	GDGW_ATTACHMENT    ListProjectTagsRequestResourceType
	CONNECT_ATTACHMENT ListProjectTagsRequestResourceType
	CFW_ATTACHMENT     ListProjectTagsRequestResourceType
	ATTACHMENTS        ListProjectTagsRequestResourceType
}

func GetListProjectTagsRequestResourceTypeEnum() ListProjectTagsRequestResourceTypeEnum {
	return ListProjectTagsRequestResourceTypeEnum{
		INSTANCE: ListProjectTagsRequestResourceType{
			value: "instance",
		},
		ROUTE_TABLE: ListProjectTagsRequestResourceType{
			value: "route-table",
		},
		VPC_ATTACHMENT: ListProjectTagsRequestResourceType{
			value: "vpc-attachment",
		},
		DGW_ATTACHMENT: ListProjectTagsRequestResourceType{
			value: "dgw-attachment",
		},
		VGW_ATTACHMENT: ListProjectTagsRequestResourceType{
			value: "vgw-attachment",
		},
		PEERING_ATTACHMENT: ListProjectTagsRequestResourceType{
			value: "peering-attachment",
		},
		VPN_ATTACHMENT: ListProjectTagsRequestResourceType{
			value: "vpn-attachment",
		},
		CAN_ATTACHMENT: ListProjectTagsRequestResourceType{
			value: "can-attachment",
		},
		ECN_ATTACHMENT: ListProjectTagsRequestResourceType{
			value: "ecn-attachment",
		},
		GDGW_ATTACHMENT: ListProjectTagsRequestResourceType{
			value: "gdgw-attachment",
		},
		CONNECT_ATTACHMENT: ListProjectTagsRequestResourceType{
			value: "connect-attachment",
		},
		CFW_ATTACHMENT: ListProjectTagsRequestResourceType{
			value: "cfw-attachment",
		},
		ATTACHMENTS: ListProjectTagsRequestResourceType{
			value: "attachments",
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
