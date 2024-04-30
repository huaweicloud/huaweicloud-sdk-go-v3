package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type StackEvent struct {

	// 资源的类型  以HCL格式的模板为例，resource_type 为 huaweicloud_vpc  ```hcl resource \"huaweicloud_vpc\" \"my_hello_world_vpc\" {   name = \"test_vpc\" } ```  以json格式的模板为例，resource_type 为 huaweicloud_vpc  ```json {   \"resource\": {     \"huaweicloud_vpc\": {       \"my_hello_world_vpc\": {         \"name\": \"test_vpc\"       }     }   } } ```
	ResourceType *string `json:"resource_type,omitempty"`

	// 资源的名称，默认为资源的逻辑名称  以HCL格式的模板为例，resource_name 为 my_hello_world_vpc  ```hcl resource \"huaweicloud_vpc\" \"my_hello_world_vpc\" {   name = \"test_vpc\" } ```  以json格式的模板为例，resource_name 为 my_hello_world_vpc  ```json {   \"resource\": {     \"huaweicloud_vpc\": {       \"my_hello_world_vpc\": {         \"name\": \"test_vpc\"       }     }   } } ```
	ResourceName *string `json:"resource_name,omitempty"`

	// 资源id的名称，即对应资源作为id使用的值的名称，当资源未创建的时候，不返回resource_id_key 此id由provider定义，因此不同的provider可能遵循了不同的命名规则，具体的命名规则请与provider开发者确认或阅读provider文档
	ResourceIdKey *string `json:"resource_id_key,omitempty"`

	// 资源id的值，即对应资源作为id使用的值，当资源未创建的时候，不返回resource_id_value
	ResourceIdValue *string `json:"resource_id_value,omitempty"`

	// 资源键，如果用户在模板中使用了count或for_each则会返回resource_key  如果用户在模板中使用count，则resource_key为从0开始的数字  以HCL格式的模板为例，模板中count为2，意味着将会生成两个资源，对应的resource_key分别为0和1  ```hcl resource \"huaweicloud_vpc\" \"my_hello_world_vpc\" {   count = 2   name = \"test_vpc\" } ```  以json格式的模板为例，模板中count为2，意味着将会生成两个资源，对应的resource_key分别为0和1  ```json {   \"resource\": {     \"huaweicloud_vpc\": {       \"my_hello_world_vpc\": {         \"name\": \"test_vpc\",         \"count\": 2       }     }   } } ```  如果用户在模板中使用for_each，则resource_key为用户自定义的字符串  以HCL格式的模板为例，resource_key分别为vpc1和vpc2  ```hcl resource \"huaweicloud_vpc\" \"my_hello_world_vpc\" {   for_each = {     \"vpc1\" = \"test_vpc\"     \"vpc2\" = \"test_vpc\"   }   name = each.value } ```  以json格式的模板为例，resource_key分别为vpc1和vpc2  ```json {   \"resource\": {     \"huaweicloud_vpc\": {       \"my_hello_world_vpc\": {         \"for_each\": {           \"vpc1\": \"test_vpc\",           \"vpc2\": \"test_vpc\"         }         \"name\": \"${each.value}\"       }     }   } } ```
	ResourceKey *string `json:"resource_key,omitempty"`

	// 事件发生的时间 格式遵循RFC3339，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z
	Time *string `json:"time,omitempty"`

	// 此次事件的类型   * `LOG` - 记录状态信息，比如当前状态，目标状态等。   * `ERROR` - 记录失败信息   * `DRIFT` - 记录资源偏移信息   * `SUMMARY` - 记录资源变更结果总结   * `CREATION_IN_PROGRESS` - 正在生成   * `CREATION_FAILED` - 生成失败   * `CREATION_COMPLETE` - 生成完成   * `DELETION_IN_PROGRESS` - 正在删除   * `DELETION_FAILED` - 删除失败   * `DELETION_COMPLETE` - 已经删除   * `UPDATE_IN_PROGRESS` - 正在更新。此处的更新特指非替换式更新，如果是替换式更新，则是DELETION后CREATION，或者CREATION后DELETION，具体以何种行为进行替换式更新由Provider定义。   * `UPDATE_FAILED` - 更新失败。此处的更新特指非替换式更新，如果是替换式更新，则是DELETION后CREATION，或者CREATION后DELETION，具体以何种行为进行替换式更新由Provider定义。   * `UPDATE_COMPLETE` - 更新完成。此处的更新特指非替换式更新，如果是替换式更新，则是DELETION后CREATION，或者CREATION后DELETION，具体以何种行为进行替换式更新由Provider定义。
	EventType *StackEventEventType `json:"event_type,omitempty"`

	// 该资源栈事件对应的详细信息
	EventMessage *string `json:"event_message,omitempty"`

	// 此事件执行所花的时间，以秒为单位
	ElapsedSeconds *int32 `json:"elapsed_seconds,omitempty"`
}

func (o StackEvent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StackEvent struct{}"
	}

	return strings.Join([]string{"StackEvent", string(data)}, " ")
}

type StackEventEventType struct {
	value string
}

type StackEventEventTypeEnum struct {
	LOG                  StackEventEventType
	ERROR                StackEventEventType
	DRIFT                StackEventEventType
	SUMMARY              StackEventEventType
	CREATION_IN_PROGRESS StackEventEventType
	CREATION_FAILED      StackEventEventType
	CREATION_COMPLETE    StackEventEventType
	DELETION_IN_PROGRESS StackEventEventType
	DELETION_FAILED      StackEventEventType
	DELETION_COMPLETE    StackEventEventType
	UPDATE_IN_PROGRESS   StackEventEventType
	UPDATE_FAILED        StackEventEventType
	UPDATE_COMPLETE      StackEventEventType
}

func GetStackEventEventTypeEnum() StackEventEventTypeEnum {
	return StackEventEventTypeEnum{
		LOG: StackEventEventType{
			value: "LOG",
		},
		ERROR: StackEventEventType{
			value: "ERROR",
		},
		DRIFT: StackEventEventType{
			value: "DRIFT",
		},
		SUMMARY: StackEventEventType{
			value: "SUMMARY",
		},
		CREATION_IN_PROGRESS: StackEventEventType{
			value: "CREATION_IN_PROGRESS",
		},
		CREATION_FAILED: StackEventEventType{
			value: "CREATION_FAILED",
		},
		CREATION_COMPLETE: StackEventEventType{
			value: "CREATION_COMPLETE",
		},
		DELETION_IN_PROGRESS: StackEventEventType{
			value: "DELETION_IN_PROGRESS",
		},
		DELETION_FAILED: StackEventEventType{
			value: "DELETION_FAILED",
		},
		DELETION_COMPLETE: StackEventEventType{
			value: "DELETION_COMPLETE",
		},
		UPDATE_IN_PROGRESS: StackEventEventType{
			value: "UPDATE_IN_PROGRESS",
		},
		UPDATE_FAILED: StackEventEventType{
			value: "UPDATE_FAILED",
		},
		UPDATE_COMPLETE: StackEventEventType{
			value: "UPDATE_COMPLETE",
		},
	}
}

func (c StackEventEventType) Value() string {
	return c.value
}

func (c StackEventEventType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StackEventEventType) UnmarshalJSON(b []byte) error {
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
