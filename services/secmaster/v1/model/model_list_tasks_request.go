package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListTasksRequest Request Object
type ListTasksRequest struct {

	// **参数解释：** 内容类型 - application/json;charset=UTF-8    普通API请求的类型  **约束限制：** 不涉及 **取值范围：** - application/json;charset=UTF-8  **默认取值：** 不涉及
	ContentType string `json:"Content-Type"`

	// **参数解释**: 工作空间ID **取值范围**: 不涉及
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**: 待办请求的偏移量 **约束限制**: 不涉及         **取值范围**: 0-9999 **默认值**:  0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 待办分页大小 **约束限制**: 不涉及         **取值范围**: 1-100 **默认值**:  10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 待办的排序字段 - create_time 按照创建时间排序 - update_time 按照更新时间排序  **约束限制**: 不涉及         **取值范围**: - create_time - update_time  **默认值**:  create_time
	SortKey *string `json:"sort_key,omitempty"`

	// **参数解释**: 待办的排序方式 - ASC 按照升序排序 - DESC 按照降序排序  **约束限制**: 不涉及         **取值范围**: - ASC - DESC  **默认值**:  DESC
	SortDir *string `json:"sort_dir,omitempty"`

	// **参数解释**: 待办的备注 **约束限制**: 不涉及         **取值范围**: 不涉及 **默认值**:  不涉及
	Note *string `json:"note,omitempty"`

	// **参数解释**: 待办的任务名称 **约束限制**: 不涉及         **取值范围**: 不涉及 **默认值**:  不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**: 待办的业务类型 - WORKFLOWPUBLISH 流程发布 - WORKFLOWNODEAPPROVE 流程节点审核 - PLAYBOOKPUBLISH 剧本发布 - PLAYBOOKNODEAPPROVE 剧本节点审核  **约束限制**: 不涉及         **取值范围**: - WORKFLOWPUBLISH - WORKFLOWNODEAPPROVE - PLAYBOOKPUBLISH - PLAYBOOKNODEAPPROVE  **默认值**:  不涉及
	BusinessType *ListTasksRequestBusinessType `json:"business_type,omitempty"`

	// **参数解释**: 待办的创建人 **约束限制**: 不涉及         **取值范围**: 不涉及 **默认值**:  不涉及
	CreatorName *string `json:"creator_name,omitempty"`

	// **参数解释**: 待办的分类 - my_todo 待处理的待办 - all_handled 已处理的待办   **约束限制**: 不涉及         **取值范围**: - my_todo - all_handled   **默认值**:  my_todo
	QueryType *ListTasksRequestQueryType `json:"query_type,omitempty"`

	// **参数解释**: 起始时间，格式是 \"yyyy-MM-dd'T'HH:mm:ss.SSS'Z'Z\" **约束限制**: 不涉及         **取值范围**: 不涉及 **默认值**:  不涉及
	FromDate *string `json:"from_date,omitempty"`

	// **参数解释**: 的截止时间，格式是 \"yyyy-MM-dd'T'HH:mm:ss.SSS'Z'Z\" **约束限制**: 不涉及         **取值范围**: 不涉及 **默认值**:  不涉及
	ToDate *string `json:"to_date,omitempty"`
}

func (o ListTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTasksRequest struct{}"
	}

	return strings.Join([]string{"ListTasksRequest", string(data)}, " ")
}

type ListTasksRequestBusinessType struct {
	value string
}

type ListTasksRequestBusinessTypeEnum struct {
	WORKFLOWPUBLISH     ListTasksRequestBusinessType
	WORKFLOWNODEAPPROVE ListTasksRequestBusinessType
	PLAYBOOKPUBLISH     ListTasksRequestBusinessType
	PLAYBOOKNODEAPPROVE ListTasksRequestBusinessType
}

func GetListTasksRequestBusinessTypeEnum() ListTasksRequestBusinessTypeEnum {
	return ListTasksRequestBusinessTypeEnum{
		WORKFLOWPUBLISH: ListTasksRequestBusinessType{
			value: "WORKFLOWPUBLISH",
		},
		WORKFLOWNODEAPPROVE: ListTasksRequestBusinessType{
			value: "WORKFLOWNODEAPPROVE",
		},
		PLAYBOOKPUBLISH: ListTasksRequestBusinessType{
			value: "PLAYBOOKPUBLISH",
		},
		PLAYBOOKNODEAPPROVE: ListTasksRequestBusinessType{
			value: "PLAYBOOKNODEAPPROVE",
		},
	}
}

func (c ListTasksRequestBusinessType) Value() string {
	return c.value
}

func (c ListTasksRequestBusinessType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTasksRequestBusinessType) UnmarshalJSON(b []byte) error {
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

type ListTasksRequestQueryType struct {
	value string
}

type ListTasksRequestQueryTypeEnum struct {
	MY_TODO     ListTasksRequestQueryType
	ALL_HANDLED ListTasksRequestQueryType
}

func GetListTasksRequestQueryTypeEnum() ListTasksRequestQueryTypeEnum {
	return ListTasksRequestQueryTypeEnum{
		MY_TODO: ListTasksRequestQueryType{
			value: "my_todo",
		},
		ALL_HANDLED: ListTasksRequestQueryType{
			value: "all_handled",
		},
	}
}

func (c ListTasksRequestQueryType) Value() string {
	return c.value
}

func (c ListTasksRequestQueryType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTasksRequestQueryType) UnmarshalJSON(b []byte) error {
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
