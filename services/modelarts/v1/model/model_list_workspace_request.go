package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListWorkspaceRequest Request Object
type ListWorkspaceRequest struct {

	// 分页列表的起始页，默认为'0'。
	Offset *int32 `json:"offset,omitempty"`

	// 指定每一页返回的最大条目数，默认为'1000'。
	Limit *int32 `json:"limit,omitempty"`

	// 指定排序字段，可选'name'、'update_time'、'status'，默认是'name'。
	SortBy *ListWorkspaceRequestSortBy `json:"sort_by,omitempty"`

	// 可选值。'asc'为递增排序。'desc'为递减排序，默认为'desc'。
	Order *ListWorkspaceRequestOrder `json:"order,omitempty"`

	// 企业项目id，指定此参数会只返回该企业项目id下的工作空间。默认显示所有工作空间。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 工作空间名称查询参数，指定此参数会模糊查询该名称的工作空间。默认显示所有工作空间。
	Name *string `json:"name,omitempty"`

	// 该参数用于筛选可访问的工作空间。指定该参数为true，则会筛选掉当前用户无权限访问的工作空间。该参数默认为false，即为显示所有工作空间。
	FilterAccessible *bool `json:"filter_accessible,omitempty"`
}

func (o ListWorkspaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkspaceRequest struct{}"
	}

	return strings.Join([]string{"ListWorkspaceRequest", string(data)}, " ")
}

type ListWorkspaceRequestSortBy struct {
	value string
}

type ListWorkspaceRequestSortByEnum struct {
	NAME        ListWorkspaceRequestSortBy
	UPDATE_TIME ListWorkspaceRequestSortBy
	STATUS      ListWorkspaceRequestSortBy
}

func GetListWorkspaceRequestSortByEnum() ListWorkspaceRequestSortByEnum {
	return ListWorkspaceRequestSortByEnum{
		NAME: ListWorkspaceRequestSortBy{
			value: "name",
		},
		UPDATE_TIME: ListWorkspaceRequestSortBy{
			value: "update_time",
		},
		STATUS: ListWorkspaceRequestSortBy{
			value: "status",
		},
	}
}

func (c ListWorkspaceRequestSortBy) Value() string {
	return c.value
}

func (c ListWorkspaceRequestSortBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListWorkspaceRequestSortBy) UnmarshalJSON(b []byte) error {
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

type ListWorkspaceRequestOrder struct {
	value string
}

type ListWorkspaceRequestOrderEnum struct {
	ASC  ListWorkspaceRequestOrder
	DESC ListWorkspaceRequestOrder
}

func GetListWorkspaceRequestOrderEnum() ListWorkspaceRequestOrderEnum {
	return ListWorkspaceRequestOrderEnum{
		ASC: ListWorkspaceRequestOrder{
			value: "asc",
		},
		DESC: ListWorkspaceRequestOrder{
			value: "desc",
		},
	}
}

func (c ListWorkspaceRequestOrder) Value() string {
	return c.value
}

func (c ListWorkspaceRequestOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListWorkspaceRequestOrder) UnmarshalJSON(b []byte) error {
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
