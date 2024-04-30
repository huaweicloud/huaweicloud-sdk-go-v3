package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StackResource 资源栈中所管理的资源信息
type StackResource struct {

	// 资源的物理id，由该资源的provider、云服务或其他服务提供方在资源部署的时候生成  注：与physical相关的参数可以在模板以外的地方，作为该资源的一种标识
	PhysicalResourceId *string `json:"physical_resource_id,omitempty"`

	// 资源的物理名称，由该资源的provider、云服务或其他服务提供方在资源部署的时候定义  注：与physical相关的参数可以在模板以外的地方，作为该资源的一种标识
	PhysicalResourceName *string `json:"physical_resource_name,omitempty"`

	// 资源的逻辑名称，由用户在模板中定义  注：与 logical 相关的参数仅仅在模板内部，作为该资源的一种标识  以HCL格式的模板为例，logical_resource_name 为 my_hello_world_vpc  ```hcl resource \"huaweicloud_vpc\" \"my_hello_world_vpc\" {   name = \"test_vpc\" } ```  以json格式的模板为例，logical_resource_name 为 my_hello_world_vpc  ```json {   \"resource\": {     \"huaweicloud_vpc\": {       \"my_hello_world_vpc\": {         \"name\": \"test_vpc\"       }     }   } } ```
	LogicalResourceName *string `json:"logical_resource_name,omitempty"`

	// 资源的类型  注：与 logical 相关的参数仅仅在模板内部，作为该资源的一种标识  以HCL格式的模板为例，logical_resource_type 为 huaweicloud_vpc  ```hcl resource \"huaweicloud_vpc\" \"my_hello_world_vpc\" {   name = \"test_vpc\" } ```  以json格式的模板为例，logical_resource_type 为 huaweicloud_vpc  ```json {   \"resource\": {     \"huaweicloud_vpc\": {       \"my_hello_world_vpc\": {         \"name\": \"test_vpc\"       }     }   } } ```
	LogicalResourceType *string `json:"logical_resource_type,omitempty"`

	// 资源的索引，如果用户在模板中使用了count或for_each则会返回index_key。如果index_key出现，则logical_resource_name + index_key可以作为该资源的一种标识  如果用户在模板中使用count，则index_key为从0开始的数字  以HCL格式的模板为例，用户在模板中可以通过`huaweicloud_vpc.my_hello_world_vpc[0]`和`huaweicloud_vpc.my_hello_world_vpc[1]`标识两个资源  ```hcl resource \"huaweicloud_vpc\" \"my_hello_world_vpc\" {   count = 2   name = \"test_vpc\" } ```  以json格式的模板为例，用户在模板中可以通过`huaweicloud_vpc.my_hello_world_vpc[0]`和`huaweicloud_vpc.my_hello_world_vpc[1]`标识两个资源  ```json {   \"resource\": {     \"huaweicloud_vpc\": {       \"my_hello_world_vpc\": {         \"name\": \"test_vpc\",         \"count\": 2       }     }   } } ```  如果用户在模板中使用for_each，则index_key为用户自定义的字符串  以HCL格式的模板为例，用户在模板中可以通过`huaweicloud_vpc.my_hello_world_vpc[\"vpc1\"]`和`huaweicloud_vpc.my_hello_world_vpc[\"vpc2\"]`标识两个资源  ```hcl resource \"huaweicloud_vpc\" \"my_hello_world_vpc\" {   for_each = {     \"vpc1\" = \"test_vpc\"     \"vpc2\" = \"test_vpc\"   }   name = each.value } ```  以json格式的模板为例，用户在模板中可以通过`huaweicloud_vpc.my_hello_world_vpc[\"vpc1\"]`和`huaweicloud_vpc.my_hello_world_vpc[\"vpc2\"]`标识两个资源  ```json {   \"resource\": {     \"huaweicloud_vpc\": {       \"my_hello_world_vpc\": {         \"for_each\": {           \"vpc1\": \"test_vpc\",           \"vpc2\": \"test_vpc\"         }         \"name\": \"${each.value}\"       }     }   } } ```
	IndexKey *string `json:"index_key,omitempty"`

	// 资源的状态 * `CREATION_IN_PROGRESS` - 正在生成 * `CREATION_FAILED`      - 生成失败 * `CREATION_COMPLETE`    - 生成完成 * `DELETION_IN_PROGRESS` - 正在删除 * `DELETION_FAILED`      - 删除失败 * `DELETION_COMPLETE`    - 已经删除 * `UPDATE_IN_PROGRESS`   - 正在更新。此处的更新特指非替换式更新，如果是替换式更新，则使用CREATION后DELETION * `UPDATE_FAILED`        - 更新失败。此处的更新特指非替换式更新，如果是替换式更新，则使用CREATION后DELETION * `UPDATE_COMPLETE`      - 更新完成。此处的更新特指非替换式更新，如果是替换式更新，则使用CREATION后DELETION
	ResourceStatus *StackResourceResourceStatus `json:"resource_status,omitempty"`

	// 当该资源状态为任意失败状态（即以 `FAILED` 结尾时），将会展示简要的错误信息总结以供debug
	StatusMessage *string `json:"status_message,omitempty"`

	// 资源属性列表
	ResourceAttributes *[]ResourceAttribute `json:"resource_attributes,omitempty"`
}

func (o StackResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StackResource struct{}"
	}

	return strings.Join([]string{"StackResource", string(data)}, " ")
}

type StackResourceResourceStatus struct {
	value string
}

type StackResourceResourceStatusEnum struct {
	CREATION_IN_PROGRESS StackResourceResourceStatus
	CREATION_FAILED      StackResourceResourceStatus
	CREATION_COMPLETE    StackResourceResourceStatus
	DELETION_IN_PROGRESS StackResourceResourceStatus
	DELETION_FAILED      StackResourceResourceStatus
	DELETION_COMPLETE    StackResourceResourceStatus
	UPDATE_IN_PROGRESS   StackResourceResourceStatus
	UPDATE_FAILED        StackResourceResourceStatus
	UPDATE_COMPLETE      StackResourceResourceStatus
}

func GetStackResourceResourceStatusEnum() StackResourceResourceStatusEnum {
	return StackResourceResourceStatusEnum{
		CREATION_IN_PROGRESS: StackResourceResourceStatus{
			value: "CREATION_IN_PROGRESS",
		},
		CREATION_FAILED: StackResourceResourceStatus{
			value: "CREATION_FAILED",
		},
		CREATION_COMPLETE: StackResourceResourceStatus{
			value: "CREATION_COMPLETE",
		},
		DELETION_IN_PROGRESS: StackResourceResourceStatus{
			value: "DELETION_IN_PROGRESS",
		},
		DELETION_FAILED: StackResourceResourceStatus{
			value: "DELETION_FAILED",
		},
		DELETION_COMPLETE: StackResourceResourceStatus{
			value: "DELETION_COMPLETE",
		},
		UPDATE_IN_PROGRESS: StackResourceResourceStatus{
			value: "UPDATE_IN_PROGRESS",
		},
		UPDATE_FAILED: StackResourceResourceStatus{
			value: "UPDATE_FAILED",
		},
		UPDATE_COMPLETE: StackResourceResourceStatus{
			value: "UPDATE_COMPLETE",
		},
	}
}

func (c StackResourceResourceStatus) Value() string {
	return c.value
}

func (c StackResourceResourceStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StackResourceResourceStatus) UnmarshalJSON(b []byte) error {
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
