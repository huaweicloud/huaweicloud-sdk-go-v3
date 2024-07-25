package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListWorkflowsRequest Request Object
type ListWorkflowsRequest struct {

	// 内容类型
	ContentType string `json:"content-type"`

	// 工作空间id
	WorkspaceId string `json:"workspace_id"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 数据量
	Limit *int32 `json:"limit,omitempty"`

	// 排序顺序，asc：升序，desc：降序
	Order *ListWorkflowsRequestOrder `json:"order,omitempty"`

	// 排序字段，create_time：创建时间，category：类型分类名称
	Sortby *ListWorkflowsRequestSortby `json:"sortby,omitempty"`

	// 是否启用
	Enabled *bool `json:"enabled,omitempty"`

	// 最新版本号
	LastVersion *bool `json:"last_version,omitempty"`

	// 流程名称
	Name *string `json:"name,omitempty"`

	// 流程描述
	Description *string `json:"description,omitempty"`

	// 数据类ID
	DataclassId *string `json:"dataclass_id,omitempty"`

	// 数据类名称
	DataclassName *string `json:"dataclass_name,omitempty"`

	// 流程类型
	AopType *string `json:"aop_type,omitempty"`
}

func (o ListWorkflowsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkflowsRequest struct{}"
	}

	return strings.Join([]string{"ListWorkflowsRequest", string(data)}, " ")
}

type ListWorkflowsRequestOrder struct {
	value string
}

type ListWorkflowsRequestOrderEnum struct {
	ASC  ListWorkflowsRequestOrder
	DESC ListWorkflowsRequestOrder
}

func GetListWorkflowsRequestOrderEnum() ListWorkflowsRequestOrderEnum {
	return ListWorkflowsRequestOrderEnum{
		ASC: ListWorkflowsRequestOrder{
			value: "asc",
		},
		DESC: ListWorkflowsRequestOrder{
			value: "desc",
		},
	}
}

func (c ListWorkflowsRequestOrder) Value() string {
	return c.value
}

func (c ListWorkflowsRequestOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListWorkflowsRequestOrder) UnmarshalJSON(b []byte) error {
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

type ListWorkflowsRequestSortby struct {
	value string
}

type ListWorkflowsRequestSortbyEnum struct {
	CATEGORY    ListWorkflowsRequestSortby
	CREATE_TIME ListWorkflowsRequestSortby
}

func GetListWorkflowsRequestSortbyEnum() ListWorkflowsRequestSortbyEnum {
	return ListWorkflowsRequestSortbyEnum{
		CATEGORY: ListWorkflowsRequestSortby{
			value: "category",
		},
		CREATE_TIME: ListWorkflowsRequestSortby{
			value: "create_time",
		},
	}
}

func (c ListWorkflowsRequestSortby) Value() string {
	return c.value
}

func (c ListWorkflowsRequestSortby) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListWorkflowsRequestSortby) UnmarshalJSON(b []byte) error {
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
