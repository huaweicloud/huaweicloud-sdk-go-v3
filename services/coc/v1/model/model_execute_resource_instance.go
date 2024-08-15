package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExecuteResourceInstance 执行脚本时需要传的机器信息
type ExecuteResourceInstance struct {

	// ecs云服务器ID
	ResourceId string `json:"resource_id"`

	// 服务器所属region
	RegionId string `json:"region_id"`

	// 资源提供者：ECS，不传默认为：ECS。请保证一次执行， 每个实例的provider是一致的。后续扩展CCE等
	Provider *string `json:"provider,omitempty"`

	// 资源提供者下资源类型，不传默认为CLOUDSERVER CLOUDSERVER:CLOUDSERVER类型 约束：  -不允许跨type支持
	Type *ExecuteResourceInstanceType `json:"type,omitempty"`

	// 支持用户自定义5个key_value形式的属性。  约束条件： - key值长度为10  - value长度为20  - map长度最大为5 - 禁止填写敏感数据
	CustomAttributes *[]Customttribute `json:"custom_attributes,omitempty"`

	// 该参数已废弃，传入该参数不会生效。
	AgentSn *string `json:"agent_sn,omitempty"`

	// 该参数已废弃，传入该参数不会生效。
	ProjectId *string `json:"project_id,omitempty"`
}

func (o ExecuteResourceInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteResourceInstance struct{}"
	}

	return strings.Join([]string{"ExecuteResourceInstance", string(data)}, " ")
}

type ExecuteResourceInstanceType struct {
	value string
}

type ExecuteResourceInstanceTypeEnum struct {
	CLOUDSERVER ExecuteResourceInstanceType
}

func GetExecuteResourceInstanceTypeEnum() ExecuteResourceInstanceTypeEnum {
	return ExecuteResourceInstanceTypeEnum{
		CLOUDSERVER: ExecuteResourceInstanceType{
			value: "CLOUDSERVER",
		},
	}
}

func (c ExecuteResourceInstanceType) Value() string {
	return c.value
}

func (c ExecuteResourceInstanceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExecuteResourceInstanceType) UnmarshalJSON(b []byte) error {
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
