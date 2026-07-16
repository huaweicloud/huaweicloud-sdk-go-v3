package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListWorkflowStepExecutionRequest Request Object
type ListWorkflowStepExecutionRequest struct {

	// 工作流的ID。
	WorkflowId string `json:"workflow_id"`

	// 返回的数据条目数。
	Limit *int32 `json:"limit,omitempty"`

	// 数据条目偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// instance order
	Order *ListWorkflowStepExecutionRequestOrder `json:"order,omitempty"`

	// 排序依据字段，例如sort_by=create_time，则表示以条目的创建时间进行排序。
	SortBy *string `json:"sort_by,omitempty"`
}

func (o ListWorkflowStepExecutionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkflowStepExecutionRequest struct{}"
	}

	return strings.Join([]string{"ListWorkflowStepExecutionRequest", string(data)}, " ")
}

type ListWorkflowStepExecutionRequestOrder struct {
	value string
}

type ListWorkflowStepExecutionRequestOrderEnum struct {
	ASC  ListWorkflowStepExecutionRequestOrder
	DESC ListWorkflowStepExecutionRequestOrder
}

func GetListWorkflowStepExecutionRequestOrderEnum() ListWorkflowStepExecutionRequestOrderEnum {
	return ListWorkflowStepExecutionRequestOrderEnum{
		ASC: ListWorkflowStepExecutionRequestOrder{
			value: "asc",
		},
		DESC: ListWorkflowStepExecutionRequestOrder{
			value: "desc",
		},
	}
}

func (c ListWorkflowStepExecutionRequestOrder) Value() string {
	return c.value
}

func (c ListWorkflowStepExecutionRequestOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListWorkflowStepExecutionRequestOrder) UnmarshalJSON(b []byte) error {
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
