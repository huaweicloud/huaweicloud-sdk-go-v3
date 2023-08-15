package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// L7ruleResp 转发规则对象
type L7ruleResp struct {

	// 转发规则ID
	Id string `json:"id"`

	// 转发规则的配置状态；该字段为预留字段，暂未启用。默认为ACTIVE。
	ProvisioningStatus string `json:"provisioning_status"`

	// 转发规则所在的项目ID。
	TenantId string `json:"tenant_id"`

	// 转发规则所在的项目ID。
	ProjectId string `json:"project_id"`

	// 转发规则的管理状态；该字段为预留字段，暂未启用。默认为true。
	AdminStateUp bool `json:"admin_state_up"`

	// 转发规则的匹配内容
	Type L7ruleRespType `json:"type"`

	// 转发规则的匹配方式。type为HOST_NAME时可以为EQUAL_TO。type为PATH时可以为REGEX， STARTS_WITH，EQUAL_TO。
	CompareType string `json:"compare_type"`

	// 是否反向匹配。使用说明：固定为false。该字段能更新但不会生效。
	Invert bool `json:"invert"`

	// 匹配内容的键值。目前匹配内容为HOST_NAME和PATH时，该字段不生效。该字段能更新但不会生效。
	Key string `json:"key"`

	// 匹配内容的值。其值不能包含空格。使用说明：当type为HOST_NAME时，取值范围：String(100)，字符串只能包含英文字母、数字、“-”或“.”，且必须以字母或数字开头。当type为PATH时，取值范围：String(128)。当转发规则的compare_type为STARTS_WITH，EQUAL_TO时，字符串只能包含英文字母、数字、^-%#&$.*+?,=!:| /()[]{}，且必须以\"/\"开头。
	Value string `json:"value"`
}

func (o L7ruleResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "L7ruleResp struct{}"
	}

	return strings.Join([]string{"L7ruleResp", string(data)}, " ")
}

type L7ruleRespType struct {
	value string
}

type L7ruleRespTypeEnum struct {
	HOST_NAME L7ruleRespType
	PATH      L7ruleRespType
}

func GetL7ruleRespTypeEnum() L7ruleRespTypeEnum {
	return L7ruleRespTypeEnum{
		HOST_NAME: L7ruleRespType{
			value: "HOST_NAME",
		},
		PATH: L7ruleRespType{
			value: "PATH",
		},
	}
}

func (c L7ruleRespType) Value() string {
	return c.value
}

func (c L7ruleRespType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *L7ruleRespType) UnmarshalJSON(b []byte) error {
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
