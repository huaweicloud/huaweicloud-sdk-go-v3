package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListWorkflowExecutionsRequest Request Object
type ListWorkflowExecutionsRequest struct {

	// 工作流的ID。
	WorkflowId string `json:"workflow_id"`

	// 工作空间ID。[获取方法请参见[查询工作空间列表](ListWorkspace.xml)。](tag:hc)未创建工作空间时默认值为“0”，存在创建并使用的工作空间，以实际取值为准。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 分页参数limit，表示单次查询的条目数上限。假如要查询20~29条记录，offset为20，limit为10。
	Limit *string `json:"limit,omitempty"`

	// 排序依据字段，例如sort_by=create_time，则表示以条目的创建时间进行排序。
	SortBy *string `json:"sort_by,omitempty"`

	// 分页参数offset，表示单次查询的条目偏移数量。假如要查询20~29条记录，offset为20，limit为10。
	Offset *string `json:"offset,omitempty"`

	// 执行记录标签。
	Labels *string `json:"labels,omitempty"`

	// 执行记录状态。
	Status *string `json:"status,omitempty"`

	// 场景ID。
	SceneId *string `json:"scene_id,omitempty"`

	// 排序的方式。该字段必须与sort_by同时使用。 缺省值: desc 枚举值： - asc：表示升序排列， - desc：降序排列。
	Order *ListWorkflowExecutionsRequestOrder `json:"order,omitempty"`
}

func (o ListWorkflowExecutionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkflowExecutionsRequest struct{}"
	}

	return strings.Join([]string{"ListWorkflowExecutionsRequest", string(data)}, " ")
}

type ListWorkflowExecutionsRequestOrder struct {
	value string
}

type ListWorkflowExecutionsRequestOrderEnum struct {
	DESC ListWorkflowExecutionsRequestOrder
	ASC  ListWorkflowExecutionsRequestOrder
}

func GetListWorkflowExecutionsRequestOrderEnum() ListWorkflowExecutionsRequestOrderEnum {
	return ListWorkflowExecutionsRequestOrderEnum{
		DESC: ListWorkflowExecutionsRequestOrder{
			value: "desc",
		},
		ASC: ListWorkflowExecutionsRequestOrder{
			value: "asc",
		},
	}
}

func (c ListWorkflowExecutionsRequestOrder) Value() string {
	return c.value
}

func (c ListWorkflowExecutionsRequestOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListWorkflowExecutionsRequestOrder) UnmarshalJSON(b []byte) error {
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
