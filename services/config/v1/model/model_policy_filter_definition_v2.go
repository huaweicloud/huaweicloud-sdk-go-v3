package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PolicyFilterDefinitionV2 规则过滤器
type PolicyFilterDefinitionV2 struct {

	// 区域ID列表
	RegionIds *[]string `json:"region_ids,omitempty"`

	// 云服务列表
	ResourceTypes *[]string `json:"resource_types,omitempty"`

	// 资源ID列表
	ResourceIds *[]string `json:"resource_ids,omitempty"`

	// 参数tags的取值逻辑，例如tags设置了a:a和b:b，当取值AND时，表示规则仅对同时绑定了a:a和b:b的资源生效；当取值为OR时，表示规则对任何绑定了a:a或b:b的资源生效。
	TagKeyLogic *PolicyFilterDefinitionV2TagKeyLogic `json:"tag_key_logic,omitempty"`

	// 生效标签列表
	Tags *[]FilterTagDetail `json:"tags,omitempty"`

	// 参数exclude_tags的取值逻辑，例如exclude_tags设置了a:a和b:b，当取值AND时，表示规则仅对同时绑定了a:a和b:b的资源生效；当取值为OR时，表示规则对任何绑定了a:a或b:b的资源生效。
	ExcludeTagKeyLogic *PolicyFilterDefinitionV2ExcludeTagKeyLogic `json:"exclude_tag_key_logic,omitempty"`

	// 排除标签列表，排除标签列表比生效标签列表有更高的优先级
	ExcludeTags *[]FilterTagDetail `json:"exclude_tags,omitempty"`
}

func (o PolicyFilterDefinitionV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyFilterDefinitionV2 struct{}"
	}

	return strings.Join([]string{"PolicyFilterDefinitionV2", string(data)}, " ")
}

type PolicyFilterDefinitionV2TagKeyLogic struct {
	value string
}

type PolicyFilterDefinitionV2TagKeyLogicEnum struct {
	AND PolicyFilterDefinitionV2TagKeyLogic
	OR  PolicyFilterDefinitionV2TagKeyLogic
}

func GetPolicyFilterDefinitionV2TagKeyLogicEnum() PolicyFilterDefinitionV2TagKeyLogicEnum {
	return PolicyFilterDefinitionV2TagKeyLogicEnum{
		AND: PolicyFilterDefinitionV2TagKeyLogic{
			value: "AND",
		},
		OR: PolicyFilterDefinitionV2TagKeyLogic{
			value: "OR",
		},
	}
}

func (c PolicyFilterDefinitionV2TagKeyLogic) Value() string {
	return c.value
}

func (c PolicyFilterDefinitionV2TagKeyLogic) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PolicyFilterDefinitionV2TagKeyLogic) UnmarshalJSON(b []byte) error {
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

type PolicyFilterDefinitionV2ExcludeTagKeyLogic struct {
	value string
}

type PolicyFilterDefinitionV2ExcludeTagKeyLogicEnum struct {
	AND PolicyFilterDefinitionV2ExcludeTagKeyLogic
	OR  PolicyFilterDefinitionV2ExcludeTagKeyLogic
}

func GetPolicyFilterDefinitionV2ExcludeTagKeyLogicEnum() PolicyFilterDefinitionV2ExcludeTagKeyLogicEnum {
	return PolicyFilterDefinitionV2ExcludeTagKeyLogicEnum{
		AND: PolicyFilterDefinitionV2ExcludeTagKeyLogic{
			value: "AND",
		},
		OR: PolicyFilterDefinitionV2ExcludeTagKeyLogic{
			value: "OR",
		},
	}
}

func (c PolicyFilterDefinitionV2ExcludeTagKeyLogic) Value() string {
	return c.value
}

func (c PolicyFilterDefinitionV2ExcludeTagKeyLogic) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PolicyFilterDefinitionV2ExcludeTagKeyLogic) UnmarshalJSON(b []byte) error {
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
