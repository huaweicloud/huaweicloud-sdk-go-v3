package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ResourceInstance struct {

	// 实例唯一id
	ResourceId string `json:"resource_id"`

	// 资源提供者：ECS。单个脚本工单， 每个实例的provider是一致的
	Provider string `json:"provider"`

	// 机器所属region的ID
	RegionId string `json:"region_id"`

	// 资源提供者下资源类型，不传默认为CLOUDSERVER CLOUDSERVER:CLOUDSERVER类型
	Type ResourceInstanceType `json:"type"`

	// 支持用户自定义5个key_value形式的属性。 约束条件：  - key值长度为20     - value长度为50     - map长度最大为5
	CustomAttributes *[]Customttribute `json:"custom_attributes,omitempty"`

	// agent纳管id。该参数即将废弃，请勿使用。
	AgentSn *string `json:"agent_sn,omitempty"`

	// agent纳管状态。该参数即将废弃，请勿使用。
	AgentStatus *string `json:"agent_status,omitempty"`

	Properties *ResourceInstanceProp `json:"properties,omitempty"`
}

func (o ResourceInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceInstance struct{}"
	}

	return strings.Join([]string{"ResourceInstance", string(data)}, " ")
}

type ResourceInstanceType struct {
	value string
}

type ResourceInstanceTypeEnum struct {
	CLOUDSERVER ResourceInstanceType
}

func GetResourceInstanceTypeEnum() ResourceInstanceTypeEnum {
	return ResourceInstanceTypeEnum{
		CLOUDSERVER: ResourceInstanceType{
			value: "CLOUDSERVER",
		},
	}
}

func (c ResourceInstanceType) Value() string {
	return c.value
}

func (c ResourceInstanceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourceInstanceType) UnmarshalJSON(b []byte) error {
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
