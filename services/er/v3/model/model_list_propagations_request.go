package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListPropagationsRequest Request Object
type ListPropagationsRequest struct {

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

	// - vpc：虚拟私有云 - vpn：vpn网关 - vgw：云专线的虚拟网关 - peering：对等连接，通过云连接CC加载不同区域的企业路由器来创建“对等连接（Peering）”连接
	ResourceType *[]ListPropagationsRequestResourceType `json:"resource_type,omitempty"`

	// 企业路由器实例状态
	State *[]ListPropagationsRequestState `json:"state,omitempty"`

	// 按关键字排序，默认按照id排序，可选值:id|name|state
	SortKey *[]string `json:"sort_key,omitempty"`

	// 返回结果按照升序或降序排列，默认为asc,降序为desc
	SortDir *[]ListPropagationsRequestSortDir `json:"sort_dir,omitempty"`
}

func (o ListPropagationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPropagationsRequest struct{}"
	}

	return strings.Join([]string{"ListPropagationsRequest", string(data)}, " ")
}

type ListPropagationsRequestResourceType struct {
	value string
}

type ListPropagationsRequestResourceTypeEnum struct {
	VPC     ListPropagationsRequestResourceType
	VPN     ListPropagationsRequestResourceType
	DGW     ListPropagationsRequestResourceType
	VGW     ListPropagationsRequestResourceType
	PEERING ListPropagationsRequestResourceType
	CAN     ListPropagationsRequestResourceType
	ECN     ListPropagationsRequestResourceType
	GDGW    ListPropagationsRequestResourceType
	CONNECT ListPropagationsRequestResourceType
	CFW     ListPropagationsRequestResourceType
}

func GetListPropagationsRequestResourceTypeEnum() ListPropagationsRequestResourceTypeEnum {
	return ListPropagationsRequestResourceTypeEnum{
		VPC: ListPropagationsRequestResourceType{
			value: "vpc",
		},
		VPN: ListPropagationsRequestResourceType{
			value: "vpn",
		},
		DGW: ListPropagationsRequestResourceType{
			value: "dgw",
		},
		VGW: ListPropagationsRequestResourceType{
			value: "vgw",
		},
		PEERING: ListPropagationsRequestResourceType{
			value: "peering",
		},
		CAN: ListPropagationsRequestResourceType{
			value: "can",
		},
		ECN: ListPropagationsRequestResourceType{
			value: "ecn",
		},
		GDGW: ListPropagationsRequestResourceType{
			value: "gdgw",
		},
		CONNECT: ListPropagationsRequestResourceType{
			value: "connect",
		},
		CFW: ListPropagationsRequestResourceType{
			value: "cfw",
		},
	}
}

func (c ListPropagationsRequestResourceType) Value() string {
	return c.value
}

func (c ListPropagationsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPropagationsRequestResourceType) UnmarshalJSON(b []byte) error {
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

type ListPropagationsRequestState struct {
	value string
}

type ListPropagationsRequestStateEnum struct {
	PENDING   ListPropagationsRequestState
	AVAILABLE ListPropagationsRequestState
	MODIFYING ListPropagationsRequestState
	DELETING  ListPropagationsRequestState
	DELETED   ListPropagationsRequestState
	FAILED    ListPropagationsRequestState
}

func GetListPropagationsRequestStateEnum() ListPropagationsRequestStateEnum {
	return ListPropagationsRequestStateEnum{
		PENDING: ListPropagationsRequestState{
			value: "pending",
		},
		AVAILABLE: ListPropagationsRequestState{
			value: "available",
		},
		MODIFYING: ListPropagationsRequestState{
			value: "modifying",
		},
		DELETING: ListPropagationsRequestState{
			value: "deleting",
		},
		DELETED: ListPropagationsRequestState{
			value: "deleted",
		},
		FAILED: ListPropagationsRequestState{
			value: "failed",
		},
	}
}

func (c ListPropagationsRequestState) Value() string {
	return c.value
}

func (c ListPropagationsRequestState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPropagationsRequestState) UnmarshalJSON(b []byte) error {
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

type ListPropagationsRequestSortDir struct {
	value string
}

type ListPropagationsRequestSortDirEnum struct {
	ASC  ListPropagationsRequestSortDir
	DESC ListPropagationsRequestSortDir
}

func GetListPropagationsRequestSortDirEnum() ListPropagationsRequestSortDirEnum {
	return ListPropagationsRequestSortDirEnum{
		ASC: ListPropagationsRequestSortDir{
			value: "asc",
		},
		DESC: ListPropagationsRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListPropagationsRequestSortDir) Value() string {
	return c.value
}

func (c ListPropagationsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPropagationsRequestSortDir) UnmarshalJSON(b []byte) error {
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
