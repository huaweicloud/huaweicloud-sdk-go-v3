package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateCentralNetworkErRouteTableAttachment 创建中心网络ER附件的请求体。
type CreateCentralNetworkErRouteTableAttachment struct {

	// 实例名字。
	Name string `json:"name"`

	// 实例描述。不支持 <>。
	Description *string `json:"description,omitempty"`

	// 企业路由器的ID。
	EnterpriseRouterId string `json:"enterprise_router_id"`

	// 企业路由器的项目ID。
	EnterpriseRouterProjectId string `json:"enterprise_router_project_id"`

	// ER路由器的regionID。
	EnterpriseRouterRegionId string `json:"enterprise_router_region_id"`

	// 中心网络平面ID。
	CentralNetworkPlaneId string `json:"central_network_plane_id"`

	// 中心网络附件对端实例的连接ID，企业路由器的连接ID或者GDGW的连接ID。
	AttachmentId *string `json:"attachment_id,omitempty"`

	// 企业路由器的路由表ID。
	EnterpriseRouterTableId string `json:"enterprise_router_table_id"`

	// 实例所属项目ID。
	AttachedErTableProjectId string `json:"attached_er_table_project_id"`

	// RegionID。
	AttachedErTableRegionId string `json:"attached_er_table_region_id"`

	// 实例ID。
	AttachedErId string `json:"attached_er_id"`

	// 实例ID。
	AttachedErTableId string `json:"attached_er_table_id"`

	// - HWCloud (华为云) - Ireland (爱尔兰)
	HostedCloud CreateCentralNetworkErRouteTableAttachmentHostedCloud `json:"hosted_cloud"`
}

func (o CreateCentralNetworkErRouteTableAttachment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCentralNetworkErRouteTableAttachment struct{}"
	}

	return strings.Join([]string{"CreateCentralNetworkErRouteTableAttachment", string(data)}, " ")
}

type CreateCentralNetworkErRouteTableAttachmentHostedCloud struct {
	value string
}

type CreateCentralNetworkErRouteTableAttachmentHostedCloudEnum struct {
	HW_CLOUD CreateCentralNetworkErRouteTableAttachmentHostedCloud
	IRELAND  CreateCentralNetworkErRouteTableAttachmentHostedCloud
}

func GetCreateCentralNetworkErRouteTableAttachmentHostedCloudEnum() CreateCentralNetworkErRouteTableAttachmentHostedCloudEnum {
	return CreateCentralNetworkErRouteTableAttachmentHostedCloudEnum{
		HW_CLOUD: CreateCentralNetworkErRouteTableAttachmentHostedCloud{
			value: "HWCloud",
		},
		IRELAND: CreateCentralNetworkErRouteTableAttachmentHostedCloud{
			value: "Ireland",
		},
	}
}

func (c CreateCentralNetworkErRouteTableAttachmentHostedCloud) Value() string {
	return c.value
}

func (c CreateCentralNetworkErRouteTableAttachmentHostedCloud) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateCentralNetworkErRouteTableAttachmentHostedCloud) UnmarshalJSON(b []byte) error {
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
