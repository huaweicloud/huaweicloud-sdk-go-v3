package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ExecutionPlanItem struct {

	// 资源的类型  以HCL格式的模板为例，resource_type 为 huaweicloud_vpc  ```hcl resource \"huaweicloud_vpc\" \"my_hello_world_vpc\" {   name = \"test_vpc\" } ```  以json格式的模板为例，resource_type 为 huaweicloud_vpc  ```json {   \"resource\": {     \"huaweicloud_vpc\": {       \"my_hello_world_vpc\": {         \"name\": \"test_vpc\"       }     }   } } ```
	ResourceType *string `json:"resource_type,omitempty"`

	// 资源的名称，默认为资源的逻辑名称  以HCL格式的模板为例，resource_name 为 my_hello_world_vpc  ```hcl resource \"huaweicloud_vpc\" \"my_hello_world_vpc\" {   name = \"test_vpc\" } ```  以json格式的模板为例，resource_name 为 my_hello_world_vpc  ```json {   \"resource\": {     \"huaweicloud_vpc\": {       \"my_hello_world_vpc\": {         \"name\": \"test_vpc\"       }     }   } } ```
	ResourceName *string `json:"resource_name,omitempty"`

	// 资源的索引，如果用户在模板中使用了count或for_each则会返回index。如果index出现，则resource_name + index可以作为该资源的一种标识  如果用户在模板中使用count，则index为从0开始的数字  以HCL格式的模板为例，用户在模板中可以通过`huaweicloud_vpc.my_hello_world_vpc[0]`和`huaweicloud_vpc.my_hello_world_vpc[1]`标识两个资源  ```hcl resource \"huaweicloud_vpc\" \"my_hello_world_vpc\" {   count = 2   name = \"test_vpc\" } ```  以json格式的模板为例，用户在模板中可以通过`huaweicloud_vpc.my_hello_world_vpc[0]`和`huaweicloud_vpc.my_hello_world_vpc[1]`标识两个资源  ```json {   \"resource\": {     \"huaweicloud_vpc\": {       \"my_hello_world_vpc\": {         \"name\": \"test_vpc\",         \"count\": 2       }     }   } } ```  如果用户在模板中使用for_each，则index为用户自定义的字符串  以HCL格式的模板为例，用户在模板中可以通过`huaweicloud_vpc.my_hello_world_vpc[\"vpc1\"]`和`huaweicloud_vpc.my_hello_world_vpc[\"vpc2\"]`标识两个资源  ```hcl resource \"huaweicloud_vpc\" \"my_hello_world_vpc\" {   for_each = {     \"vpc1\" = \"test_vpc\"     \"vpc2\" = \"test_vpc\"   }   name = each.value } ```  以json格式的模板为例，用户在模板中可以通过`huaweicloud_vpc.my_hello_world_vpc[\"vpc1\"]`和`huaweicloud_vpc.my_hello_world_vpc[\"vpc2\"]`标识两个资源  ```json {   \"resource\": {     \"huaweicloud_vpc\": {       \"my_hello_world_vpc\": {         \"for_each\": {           \"vpc1\": \"test_vpc\",           \"vpc2\": \"test_vpc\"         }         \"name\": \"${each.value}\"       }     }   } } ```
	Index *string `json:"index,omitempty"`

	// 资源变更的类型   * `ADD` - 新增资源   * `ADD_THEN_DELETE` - 由不可更新的资源返回，先创建新资源，再删除旧资源   * `DELETE ` - 删除资源   * `DELETE_THEN_ADD` - 由不可更新的资源返回，先删除旧资源，再创建新资源   * `UPDATE` - 更新资源    * `NO_OPERATION` - 仅变更资源的依赖关系，但是对资源本身并无修改的操作
	Action *ExecutionPlanItemAction `json:"action,omitempty"`

	// 触发该项目变更的原因，例如用户更新模板；远端删除资源等等
	ActionReason *string `json:"action_reason,omitempty"`

	// 该项目所属的provider名称。
	ProviderName *string `json:"provider_name,omitempty"`

	// 资源模式   * `DATA` - 指可以在模板解析期间运行和获取服务端数据的资源类型，不会操作基础设施组件   * `RESOURCE` - 指通过模板管理的由服务定义的基础设施组件抽象，可以是物理资源也可以是逻辑资源
	Mode *ExecutionPlanItemMode `json:"mode,omitempty"`

	// 当前资源的变更是否由偏差导致。  偏差，也叫漂移。指的是资源被资源编排服务创建以后，又经历过非资源编排服务触发的修改，如手动修改、调用SDK修改等，使得资源的配置与本服务所记录的资源的配置不一致。这种不一致便称为偏差。  当资源产生偏差以后： * 如果用户试图创建执行计划，则会提示用户产生偏差 * 如果用户直接部署，则偏差有可能被覆盖，资源编排服务只保证资源和模板最终一致。  资源的偏差有两种类型： * 资源定位属性被修改：如果是定位属性被修改，常见于删除后重建，此时资源已经不属于同一个资源。资源编排服务会认为此资源已经被删除，会尝试创建一个新的资源。 * 资源普通属性被修改：如果是普通属性被修改，则资源编排服务依然可以找到资源，但是下次部署会尝试修复偏差，即将资源保持和模板最终一致。  **注：资源编排服务团队极力推荐，如果资源是通过本服务创建的，请一直使用本服务进行维护和更新以确保资源和模板保持一致。建议非紧急事件以外的情况不要手动调整。**
	Drifted *bool `json:"drifted,omitempty"`

	// 当前资源的变更是否是导入的。
	Imported *bool `json:"imported,omitempty"`

	// 资源的物理id，是唯一id，由为该资源提供服务的provider、云服务或其他服务提供方在资源部署的时候生成
	ResourceId *string `json:"resource_id,omitempty"`

	// 执行计划项目中变更的属性，当无属性变更时为空列表。
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
