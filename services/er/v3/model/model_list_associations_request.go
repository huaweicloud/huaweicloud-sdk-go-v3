package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAssociationsRequest Request Object
type ListAssociationsRequest struct {

	// 企业路由器实例ID
	ErId string `json:"er_id"`

	// 路由表ID
	RouteTableId string `json:"route_table_id"`

	// 每页返回的个数。 取值范围：0~2000。
	Limit *int32 `json:"limit,omitempty"`

	// 上一页最后一条记录的企业路由器实例的id，为空时为查询第一页。 使用说明：必须与limit一起使用。
	Marker *string `json:"marker,omitempty"`

	// 连接ID
	AttachmentId *[]string `json:"attachment_id,omitempty"`

	// - vpc：虚拟私有云 - vpn：vpn网关 - vgw：云专线的虚拟网关 - peering：对等连接，通过云连接CC加载不同区域的企业路由器来创建“对等连接（Peering）”连接 -  -  -  -
	ResourceType *[]ListAssociationsRequestResourceType `json:"resource_type,omitempty"`

	// 状态
	State *[]ListAssociationsRequestState `json:"state,omitempty"`

	// 按关键字排序，默认按照id排序，可选值:id|name|state
	SortKey *[]string `json:"sort_key,omitempty"`

	// 返回结果按照升序或降序排列，默认为asc,降序为desc
	SortDir *[]ListAssociationsRequestSortDir `json:"sort_dir,omitempty"`
}

func (o ListAssociationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssociationsRequest struct{}"
	}

	return strings.Join([]string{"ListAssociationsRequest", string(data)}, " ")
}

type ListAssociationsRequestResourceType struct {
	value string
}

type ListAssociationsRequestResourceTypeEnum struct {
	VPC     ListAssociationsRequestResourceType
	VPN     ListAssociationsRequestResourceType
	DGW     ListAssociationsRequestResourceType
	VGW     ListAssociationsRequestResourceType
	PEERING ListAssociationsRequestResourceType
	CAN     ListAssociationsRequestResourceType
	ECN     ListAssociationsRequestResourceType
	GDGW    ListAssociationsRequestResourceType
	CONNECT ListAssociationsRequestResourceType
	CFW     ListAssociationsRequestResourceType
}

func GetListAssociationsRequestResourceTypeEnum() ListAssociationsRequestResourceTypeEnum {
	return ListAssociationsRequestResourceTypeEnum{
		VPC: ListAssociationsRequestResourceType{
			value: "vpc",
		},
		VPN: ListAssociationsRequestResourceType{
			value: "vpn",
		},
		DGW: ListAssociationsRequestResourceType{
			value: "dgw",
		},
		VGW: ListAssociationsRequestResourceType{
			value: "vgw",
		},
		PEERING: ListAssociationsRequestResourceType{
			value: "peering",
		},
		CAN: ListAssociationsRequestResourceType{
			value: "can",
		},
		ECN: ListAssociationsRequestResourceType{
			value: "ecn",
		},
		GDGW: ListAssociationsRequestResourceType{
			value: "gdgw",
		},
		CONNECT: ListAssociationsRequestResourceType{
			value: "connect",
		},
		CFW: ListAssociationsRequestResourceType{
			value: "cfw",
		},
	}
}

func (c ListAssociationsRequestResourceType) Value() string {
	return c.value
}

func (c ListAssociationsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAssociationsRequestResourceType) UnmarshalJSON(b []byte) error {
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

type ListAssociationsRequestState struct {
	value string
}

type ListAssociationsRequestStateEnum struct {
	PENDING   ListAssociationsRequestState
	AVAILABLE ListAssociationsRequestState
	DELETING  ListAssociationsRequestState
	DELETED   ListAssociationsRequestState
	FAILED    ListAssociationsRequestState
	MODIFYING ListAssociationsRequestState
}

func GetListAssociationsRequestStateEnum() ListAssociationsRequestStateEnum {
	return ListAssociationsRequestStateEnum{
		PENDING: ListAssociationsRequestState{
			value: "pending",
		},
		AVAILABLE: ListAssociationsRequestState{
			value: "available",
		},
		DELETING: ListAssociationsRequestState{
			value: "deleting",
		},
		DELETED: ListAssociationsRequestState{
			value: "deleted",
		},
		FAILED: ListAssociationsRequestState{
			value: "failed",
		},
		MODIFYING: ListAssociationsRequestState{
			value: "modifying",
		},
	}
}

func (c ListAssociationsRequestState) Value() string {
	return c.value
}

func (c ListAssociationsRequestState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAssociationsRequestState) UnmarshalJSON(b []byte) error {
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

type ListAssociationsRequestSortDir struct {
	value string
}

type ListAssociationsRequestSortDirEnum struct {
	ASC  ListAssociationsRequestSortDir
	DESC ListAssociationsRequestSortDir
}

func GetListAssociationsRequestSortDirEnum() ListAssociationsRequestSortDirEnum {
	return ListAssociationsRequestSortDirEnum{
		ASC: ListAssociationsRequestSortDir{
			value: "asc",
		},
		DESC: ListAssociationsRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListAssociationsRequestSortDir) Value() string {
	return c.value
}

func (c ListAssociationsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAssociationsRequestSortDir) UnmarshalJSON(b []byte) error {
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
