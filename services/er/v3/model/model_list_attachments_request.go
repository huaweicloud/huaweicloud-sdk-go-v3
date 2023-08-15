package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAttachmentsRequest Request Object
type ListAttachmentsRequest struct {

	// 企业路由器实例ID
	ErId string `json:"er_id"`

	// 每页返回的个数。 取值范围：0~2000。
	Limit *int32 `json:"limit,omitempty"`

	// 上一页最后一条记录的企业路由器实例的id，为空时为查询第一页。 使用说明：必须与limit一起使用。
	Marker *string `json:"marker,omitempty"`

	// 连接状态:pending|available|modifying|deleting|deleted|failed|pending_acceptance|rejected|initiating_request
	State *[]ListAttachmentsRequestState `json:"state,omitempty"`

	// - vpc：虚拟私有云 - vpn：vpn网关 - vgw：云专线的虚拟网关 - peering：对等连接，通过云连接CC加载不同区域的企业路由器来创建“对等连接（Peering）”连接 -  -  -  -
	ResourceType *[]ListAttachmentsRequestResourceType `json:"resource_type,omitempty"`

	// 连接对应的资源ID列表
	ResourceId *[]string `json:"resource_id,omitempty"`

	// 按关键字排序，默认按照id排序，可选值:id|name|state
	SortKey *[]string `json:"sort_key,omitempty"`

	// 返回结果按照升序或降序排列，默认为asc,降序为desc
	SortDir *[]ListAttachmentsRequestSortDir `json:"sort_dir,omitempty"`
}

func (o ListAttachmentsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAttachmentsRequest struct{}"
	}

	return strings.Join([]string{"ListAttachmentsRequest", string(data)}, " ")
}

type ListAttachmentsRequestState struct {
	value string
}

type ListAttachmentsRequestStateEnum struct {
	PENDING            ListAttachmentsRequestState
	AVAILABLE          ListAttachmentsRequestState
	MODIFYING          ListAttachmentsRequestState
	DELETING           ListAttachmentsRequestState
	DELETED            ListAttachmentsRequestState
	FAILED             ListAttachmentsRequestState
	INITIATING_REQUEST ListAttachmentsRequestState
	REJECTED           ListAttachmentsRequestState
	PENDING_ACCEPTANCE ListAttachmentsRequestState
}

func GetListAttachmentsRequestStateEnum() ListAttachmentsRequestStateEnum {
	return ListAttachmentsRequestStateEnum{
		PENDING: ListAttachmentsRequestState{
			value: "pending",
		},
		AVAILABLE: ListAttachmentsRequestState{
			value: "available",
		},
		MODIFYING: ListAttachmentsRequestState{
			value: "modifying",
		},
		DELETING: ListAttachmentsRequestState{
			value: "deleting",
		},
		DELETED: ListAttachmentsRequestState{
			value: "deleted",
		},
		FAILED: ListAttachmentsRequestState{
			value: "failed",
		},
		INITIATING_REQUEST: ListAttachmentsRequestState{
			value: "initiating_request",
		},
		REJECTED: ListAttachmentsRequestState{
			value: "rejected",
		},
		PENDING_ACCEPTANCE: ListAttachmentsRequestState{
			value: "pending_acceptance",
		},
	}
}

func (c ListAttachmentsRequestState) Value() string {
	return c.value
}

func (c ListAttachmentsRequestState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAttachmentsRequestState) UnmarshalJSON(b []byte) error {
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

type ListAttachmentsRequestResourceType struct {
	value string
}

type ListAttachmentsRequestResourceTypeEnum struct {
	VPC     ListAttachmentsRequestResourceType
	VPN     ListAttachmentsRequestResourceType
	DGW     ListAttachmentsRequestResourceType
	VGW     ListAttachmentsRequestResourceType
	PEERING ListAttachmentsRequestResourceType
	CAN     ListAttachmentsRequestResourceType
	ECN     ListAttachmentsRequestResourceType
	GDGW    ListAttachmentsRequestResourceType
	CONNECT ListAttachmentsRequestResourceType
	CFW     ListAttachmentsRequestResourceType
}

func GetListAttachmentsRequestResourceTypeEnum() ListAttachmentsRequestResourceTypeEnum {
	return ListAttachmentsRequestResourceTypeEnum{
		VPC: ListAttachmentsRequestResourceType{
			value: "vpc",
		},
		VPN: ListAttachmentsRequestResourceType{
			value: "vpn",
		},
		DGW: ListAttachmentsRequestResourceType{
			value: "dgw",
		},
		VGW: ListAttachmentsRequestResourceType{
			value: "vgw",
		},
		PEERING: ListAttachmentsRequestResourceType{
			value: "peering",
		},
		CAN: ListAttachmentsRequestResourceType{
			value: "can",
		},
		ECN: ListAttachmentsRequestResourceType{
			value: "ecn",
		},
		GDGW: ListAttachmentsRequestResourceType{
			value: "gdgw",
		},
		CONNECT: ListAttachmentsRequestResourceType{
			value: "connect",
		},
		CFW: ListAttachmentsRequestResourceType{
			value: "cfw",
		},
	}
}

func (c ListAttachmentsRequestResourceType) Value() string {
	return c.value
}

func (c ListAttachmentsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAttachmentsRequestResourceType) UnmarshalJSON(b []byte) error {
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

type ListAttachmentsRequestSortDir struct {
	value string
}

type ListAttachmentsRequestSortDirEnum struct {
	ASC  ListAttachmentsRequestSortDir
	DESC ListAttachmentsRequestSortDir
}

func GetListAttachmentsRequestSortDirEnum() ListAttachmentsRequestSortDirEnum {
	return ListAttachmentsRequestSortDirEnum{
		ASC: ListAttachmentsRequestSortDir{
			value: "asc",
		},
		DESC: ListAttachmentsRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListAttachmentsRequestSortDir) Value() string {
	return c.value
}

func (c ListAttachmentsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAttachmentsRequestSortDir) UnmarshalJSON(b []byte) error {
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
