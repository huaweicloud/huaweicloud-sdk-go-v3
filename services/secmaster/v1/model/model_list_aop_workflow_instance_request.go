package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAopWorkflowInstanceRequest Request Object
type ListAopWorkflowInstanceRequest struct {

	// **参数解释：** 内容类型 - application/json;charset=UTF-8    普通API请求的类型  **约束限制：** 不涉及 **取值范围：** - application/json;charset=UTF-8  **默认取值：** 不涉及
	ContentType string `json:"Content-Type"`

	// **参数解释**: 工作空间ID **取值范围**: 不涉及
	WorkspaceId string `json:"workspace_id"`

	// **参数解释：** 偏移量 **约束限制：** 0-10000 **取值范围：** 不涉及 **默认取值：** 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 数据量 **约束限制：** 1-100 **取值范围：** 不涉及 **默认取值：** 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 排序字段， - create_time：创建时间 - update_time：更新时间  **约束限制：** 不涉及 **取值范围：** - create_time - update_time  **默认取值：** create_time
	SortKey *ListAopWorkflowInstanceRequestSortKey `json:"sort_key,omitempty"`

	// **参数解释：** 排序顺序 - ASC：升序 - DESC：降序  **约束限制：** 不涉及 **取值范围：** - ASC：升序 - DESC：降序  **默认取值：** DESC
	SortDir *string `json:"sort_dir,omitempty"`

	// **参数解释**: 开始时间 **约束限制**: 不涉及               **取值范围**: 不涉及 **默认值**:  不涉及
	FromDate *string `json:"from_date,omitempty"`

	// **参数解释**: 截止时间 **约束限制**: 不涉及               **取值范围**: 不涉及 **默认值**:  不涉及
	ToDate *string `json:"to_date,omitempty"`

	// **参数解释**: 流程的ID **约束限制**: 不涉及               **取值范围**: 不涉及 **默认值**:  不涉及
	WorkflowId *string `json:"workflow_id,omitempty"`

	// **参数解释**: 实例的ID **约束限制**: 不涉及               **取值范围**: 不涉及 **默认值**:  不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**: 实例的名字 **约束限制**: 不涉及               **取值范围**: 不涉及 **默认值**:  不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**: 触发流程对象的id **约束限制**: 不涉及               **取值范围**: 不涉及 **默认值**:  不涉及
	DataobjectId *string `json:"dataobject_id,omitempty"`

	// **参数解释**: 数据类的ID **约束限制**: 不涉及               **取值范围**: 不涉及 **默认值**:  不涉及
	DataclassId *string `json:"dataclass_id,omitempty"`

	// **参数解释**: 剧本ID **约束限制**: 不涉及               **取值范围**: 不涉及 **默认值**:  不涉及
	PlaybookId *string `json:"playbook_id,omitempty"`

	// **参数解释**: 防线ID **约束限制**: 不涉及               **取值范围**: 不涉及 **默认值**:  不涉及
	DefenceId *string `json:"defence_id,omitempty"`

	// **参数解释**:          流程实例的状态 - RUNNING   运行中 - FAILED    运行失败 - FINISHED  运行结束 - RETRYING  重试中 - TERMINATING 终止中 - TERMINATED  已终止  **约束限制**: 不涉及               **取值范围**: - RUNNING    - FAILED     - FINISHED   - RETRYING   - TERMINATING  - TERMINATED  **默认值**:  不涉及
	Status *string `json:"status,omitempty"`

	// **参数解释**:          触发方式 - DEBUG   调试触发 - TIMER   定时触发 - EVENT   事件触发 - MANUAL  手动触发  **约束限制**: 不涉及               **取值范围**: - DEBUG - TIMER - EVENT - MANUAL  **默认值**:  不涉及
	TriggerType *string `json:"trigger_type,omitempty"`
}

func (o ListAopWorkflowInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAopWorkflowInstanceRequest struct{}"
	}

	return strings.Join([]string{"ListAopWorkflowInstanceRequest", string(data)}, " ")
}

type ListAopWorkflowInstanceRequestSortKey struct {
	value string
}

type ListAopWorkflowInstanceRequestSortKeyEnum struct {
	UPDATE_TIME ListAopWorkflowInstanceRequestSortKey
	CREATE_TIME ListAopWorkflowInstanceRequestSortKey
}

func GetListAopWorkflowInstanceRequestSortKeyEnum() ListAopWorkflowInstanceRequestSortKeyEnum {
	return ListAopWorkflowInstanceRequestSortKeyEnum{
		UPDATE_TIME: ListAopWorkflowInstanceRequestSortKey{
			value: "update_time",
		},
		CREATE_TIME: ListAopWorkflowInstanceRequestSortKey{
			value: "create_time",
		},
	}
}

func (c ListAopWorkflowInstanceRequestSortKey) Value() string {
	return c.value
}

func (c ListAopWorkflowInstanceRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAopWorkflowInstanceRequestSortKey) UnmarshalJSON(b []byte) error {
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
