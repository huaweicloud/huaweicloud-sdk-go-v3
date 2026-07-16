package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListWorkflowsRequest Request Object
type ListWorkflowsRequest struct {

	// 工作流名称。
	Name *string `json:"name,omitempty"`

	// 工作流描述信息。
	Description *string `json:"description,omitempty"`

	// 工作流状态。
	Status *string `json:"status,omitempty"`

	// 工作流标签。
	Labels *string `json:"labels,omitempty"`

	// 工作流模板ID。
	TemplateId *string `json:"template_id,omitempty"`

	// 分页参数limit，表示单次查询的条目数上限。假如要查询20~29条记录，offset为20，limit为10。
	Limit *string `json:"limit,omitempty"`

	// 分页参数offset，表示单次查询的条目偏移数量。假如要查询20~29条记录，offset为20，limit为10。
	Offset *string `json:"offset,omitempty"`

	// 排序依据字段，例如sort_by=create_time，则表示以条目的创建时间进行排序。
	SortBy *string `json:"sort_by,omitempty"`

	// 过滤方式。可选值如下： - equal表示精确匹配。 - contain表示模糊匹配。  具体过滤的字段，由各个接口额外定义参数。例如Workflow支持按照名称（name）进行过滤，则相应的过滤字段为name。name=workflow&search_type=contain表示查询名称中含有Workflow字样的所有工作流。
	SearchType *ListWorkflowsRequestSearchType `json:"search_type,omitempty"`
}

func (o ListWorkflowsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkflowsRequest struct{}"
	}

	return strings.Join([]string{"ListWorkflowsRequest", string(data)}, " ")
}

type ListWorkflowsRequestSearchType struct {
	value string
}

type ListWorkflowsRequestSearchTypeEnum struct {
	CONTAIN ListWorkflowsRequestSearchType
	EQUAL   ListWorkflowsRequestSearchType
}

func GetListWorkflowsRequestSearchTypeEnum() ListWorkflowsRequestSearchTypeEnum {
	return ListWorkflowsRequestSearchTypeEnum{
		CONTAIN: ListWorkflowsRequestSearchType{
			value: "contain",
		},
		EQUAL: ListWorkflowsRequestSearchType{
			value: "equal",
		},
	}
}

func (c ListWorkflowsRequestSearchType) Value() string {
	return c.value
}

func (c ListWorkflowsRequestSearchType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListWorkflowsRequestSearchType) UnmarshalJSON(b []byte) error {
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
