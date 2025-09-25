package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSqlPatchRequest Request Object
type ShowSqlPatchRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us **默认取值**: en-us
	XLanguage *ShowSqlPatchRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释** : 节点ID。 **约束限制** : 不涉及。 **取值范围** : 由 32 个字符（英文字母、数字、-或\\），后跟 \"no\"，再跟 \"14\" 或 \"20\" 组成的字符串。 **默认取值** : 不涉及。
	NodeId string `json:"node_id"`

	// **参数解释** : 慢SQL的ID。 **约束限制** : 不涉及。 **取值范围** : 由数字组成，且长度为1~256个字符。 **默认取值** : 不涉及。
	SqlId string `json:"sql_id"`

	// **参数解释** : 数据库名称。慢SQL场景可以不填，其他场景必填。 **约束限制** : 不能使用模板库，且是已存在的数据库名称。 模板库包括 template0 ，template1。 **取值范围** : 只能由字母、数字、_及$组成，且长度为0~63个字符，不能以数字开头。 **默认取值** : 不涉及。
	DatabaseName *string `json:"database_name,omitempty"`
}

func (o ShowSqlPatchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlPatchRequest struct{}"
	}

	return strings.Join([]string{"ShowSqlPatchRequest", string(data)}, " ")
}

type ShowSqlPatchRequestXLanguage struct {
	value string
}

type ShowSqlPatchRequestXLanguageEnum struct {
	ZH_CN ShowSqlPatchRequestXLanguage
	EN_US ShowSqlPatchRequestXLanguage
}

func GetShowSqlPatchRequestXLanguageEnum() ShowSqlPatchRequestXLanguageEnum {
	return ShowSqlPatchRequestXLanguageEnum{
		ZH_CN: ShowSqlPatchRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowSqlPatchRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowSqlPatchRequestXLanguage) Value() string {
	return c.value
}

func (c ShowSqlPatchRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSqlPatchRequestXLanguage) UnmarshalJSON(b []byte) error {
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
