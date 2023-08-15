package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListRouteTablesRequest Request Object
type ListRouteTablesRequest struct {

	// 企业路由器实例ID
	ErId string `json:"er_id"`

	// 每页返回的个数。 取值范围：0~2000。
	Limit *int32 `json:"limit,omitempty"`

	// 上一页最后一条记录的企业路由器实例的id，为空时为查询第一页。 使用说明：必须与limit一起使用。
	Marker *string `json:"marker,omitempty"`

	// 状态
	State *[]ListRouteTablesRequestState `json:"state,omitempty"`

	// 是否为默认传播路由表
	IsDefaultPropagationTable *bool `json:"is_default_propagation_table,omitempty"`

	// 是否为默认关联路由表
	IsDefaultAssociationTable *bool `json:"is_default_association_table,omitempty"`

	// 按关键字排序，默认按照id排序，可选值:id|name|state
	SortKey *[]string `json:"sort_key,omitempty"`

	// 返回结果按照升序或降序排列，默认为asc,降序为desc
	SortDir *[]ListRouteTablesRequestSortDir `json:"sort_dir,omitempty"`
}

func (o ListRouteTablesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRouteTablesRequest struct{}"
	}

	return strings.Join([]string{"ListRouteTablesRequest", string(data)}, " ")
}

type ListRouteTablesRequestState struct {
	value string
}

type ListRouteTablesRequestStateEnum struct {
	PENDING   ListRouteTablesRequestState
	AVAILABLE ListRouteTablesRequestState
	DELETING  ListRouteTablesRequestState
	DELETED   ListRouteTablesRequestState
	FAILED    ListRouteTablesRequestState
}

func GetListRouteTablesRequestStateEnum() ListRouteTablesRequestStateEnum {
	return ListRouteTablesRequestStateEnum{
		PENDING: ListRouteTablesRequestState{
			value: "pending",
		},
		AVAILABLE: ListRouteTablesRequestState{
			value: "available",
		},
		DELETING: ListRouteTablesRequestState{
			value: "deleting",
		},
		DELETED: ListRouteTablesRequestState{
			value: "deleted",
		},
		FAILED: ListRouteTablesRequestState{
			value: "failed",
		},
	}
}

func (c ListRouteTablesRequestState) Value() string {
	return c.value
}

func (c ListRouteTablesRequestState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRouteTablesRequestState) UnmarshalJSON(b []byte) error {
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

type ListRouteTablesRequestSortDir struct {
	value string
}

type ListRouteTablesRequestSortDirEnum struct {
	ASC  ListRouteTablesRequestSortDir
	DESC ListRouteTablesRequestSortDir
}

func GetListRouteTablesRequestSortDirEnum() ListRouteTablesRequestSortDirEnum {
	return ListRouteTablesRequestSortDirEnum{
		ASC: ListRouteTablesRequestSortDir{
			value: "asc",
		},
		DESC: ListRouteTablesRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListRouteTablesRequestSortDir) Value() string {
	return c.value
}

func (c ListRouteTablesRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRouteTablesRequestSortDir) UnmarshalJSON(b []byte) error {
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
