package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 执行计划元素，承载执行计划中变更的细节。
type ExecutionPlanItem struct {

	// 资源变更的类型，这里，IN_PLACE_UPDATE、ADD_THEN_DELETE和 DELETE_THEN_ADD均为更新操作，IN_PLACE_UPDATE指原地更新； 而对于不可更新的资源，ADD_THEN_DELETE是先创建新的，再删除旧的；DELETE_THEN_ADD是先删除旧的，再创建新的. 执行计划的执行状态，只有当AVAILABLE的时候才可以使用apply执行 * `ADD` - 新建资源 * `ADD_THEN_DELETE` - 对于不可更新的资源执行先创建再删除的操作 * `DELETE ` - 删除资源 * `DELETE_THEN_ADD` - 对于不可更新的资源执行先删除在创建的操作 * `UPDATE` - 更新资源  * `IN_PLACE_UPDATE` - 更新资源的操作 * `NO_OPERATION` - 变更资源的依赖关系，但是对资源本身并无修改的操作
	Action *ExecutionPlanItemAction `json:"action,omitempty"`

	// 表示该动作触发的原因，例如用户更新模板；远端删除资源等等
	ActionReason *string `json:"action_reason,omitempty"`

	// 表示该资源所属的provider名字。
	ProviderName *string `json:"provider_name,omitempty"`

	// 当前资源在HCL模板中对应的类型。
	ResourceType *string `json:"resource_type,omitempty"`

	// 当前资源的在HCL模板中指定的名字。
	ResourceName *string `json:"resource_name,omitempty"`

	// 表示资源对应的index，例如对于使用count构建的资源，其类型和名字一样，但是index是从1到count的数值；对于使用for_each创建的资源，index可以是for_each中指定的key名。
	Index *string `json:"index,omitempty"`

	// * `DATA` - 指可以在模板解析期间运行和获取服务端数据的资源类型，不会操作基础设施组件 * `RESOURCE` - 指通过模板管理的由服务定义的基础设施组件抽象，可以是物理资源也可以是逻辑资源
	Mode *ExecutionPlanItemMode `json:"mode,omitempty"`

	// 当前资源的变更是否由配置漂移导致。
	Drifted *bool `json:"drifted,omitempty"`

	// 当前资源的唯一ID，当操作类型为创建时为空。当资源为新建时为空。注意resouce_name是资源在HCL模板中定义的名字，resource_id是provider提供的唯一ID。
	ResourceId *string `json:"resource_id,omitempty"`

	// 执行计划元素变更的属性，当无属性变更时为空。
	Attributes *[]ExecutionPlanDiffAttribute `json:"attributes,omitempty"`
}

func (o ExecutionPlanItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecutionPlanItem struct{}"
	}

	return strings.Join([]string{"ExecutionPlanItem", string(data)}, " ")
}

type ExecutionPlanItemAction struct {
	value string
}

type ExecutionPlanItemActionEnum struct {
	ADD             ExecutionPlanItemAction
	ADD_THEN_DELETE ExecutionPlanItemAction
	DELETE          ExecutionPlanItemAction
	DELETE_THEN_ADD ExecutionPlanItemAction
	UPDATE          ExecutionPlanItemAction
	IN_PLACE_UPDATE ExecutionPlanItemAction
	NO_OPERATION    ExecutionPlanItemAction
}

func GetExecutionPlanItemActionEnum() ExecutionPlanItemActionEnum {
	return ExecutionPlanItemActionEnum{
		ADD: ExecutionPlanItemAction{
			value: "ADD",
		},
		ADD_THEN_DELETE: ExecutionPlanItemAction{
			value: "ADD_THEN_DELETE",
		},
		DELETE: ExecutionPlanItemAction{
			value: "DELETE",
		},
		DELETE_THEN_ADD: ExecutionPlanItemAction{
			value: "DELETE_THEN_ADD",
		},
		UPDATE: ExecutionPlanItemAction{
			value: "UPDATE",
		},
		IN_PLACE_UPDATE: ExecutionPlanItemAction{
			value: "IN_PLACE_UPDATE",
		},
		NO_OPERATION: ExecutionPlanItemAction{
			value: "NO_OPERATION",
		},
	}
}

func (c ExecutionPlanItemAction) Value() string {
	return c.value
}

func (c ExecutionPlanItemAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExecutionPlanItemAction) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ExecutionPlanItemMode struct {
	value string
}

type ExecutionPlanItemModeEnum struct {
	DATA     ExecutionPlanItemMode
	RESOURCE ExecutionPlanItemMode
}

func GetExecutionPlanItemModeEnum() ExecutionPlanItemModeEnum {
	return ExecutionPlanItemModeEnum{
		DATA: ExecutionPlanItemMode{
			value: "DATA",
		},
		RESOURCE: ExecutionPlanItemMode{
			value: "RESOURCE",
		},
	}
}

func (c ExecutionPlanItemMode) Value() string {
	return c.value
}

func (c ExecutionPlanItemMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExecutionPlanItemMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
