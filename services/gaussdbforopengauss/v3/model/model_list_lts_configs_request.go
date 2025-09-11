package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListLtsConfigsRequest Request Object
type ListLtsConfigsRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *ListLtsConfigsRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**: 实例名称。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceName *string `json:"instance_name,omitempty"`

	// **参数解释**: 实例类型，不传该参数默认查询全部实例类型。 **约束限制**: 不涉及。 **取值范围**: - Ha 为集中式类型。 - Independent 为独立部署类型。 - Combined 为混合部署类型。  **默认取值**: 不涉及。
	InstanceMode *ListLtsConfigsRequestInstanceMode `json:"instance_mode,omitempty"`

	// **参数解释**: 企业项目ID。只有企业租户时该参数才生效。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询。例如：该参数指定为1，limit指定为10，则只展示第2-11条数据。 **约束限制**: 不涉及。 **取值范围**: [0, 2^31-1] **默认取值**: 默认为0（偏移0条数据，表示从第一条数据开始查询）。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 查询记录数。例如该参数设定为10，则查询结果最多只显示10条记录。 **约束限制**: 不涉及。 **取值范围**: [1, 50] **默认取值**: 50
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListLtsConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLtsConfigsRequest struct{}"
	}

	return strings.Join([]string{"ListLtsConfigsRequest", string(data)}, " ")
}

type ListLtsConfigsRequestXLanguage struct {
	value string
}

type ListLtsConfigsRequestXLanguageEnum struct {
	ZH_CN ListLtsConfigsRequestXLanguage
	EN_US ListLtsConfigsRequestXLanguage
}

func GetListLtsConfigsRequestXLanguageEnum() ListLtsConfigsRequestXLanguageEnum {
	return ListLtsConfigsRequestXLanguageEnum{
		ZH_CN: ListLtsConfigsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListLtsConfigsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListLtsConfigsRequestXLanguage) Value() string {
	return c.value
}

func (c ListLtsConfigsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListLtsConfigsRequestXLanguage) UnmarshalJSON(b []byte) error {
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

type ListLtsConfigsRequestInstanceMode struct {
	value string
}

type ListLtsConfigsRequestInstanceModeEnum struct {
	HA          ListLtsConfigsRequestInstanceMode
	INDEPENDENT ListLtsConfigsRequestInstanceMode
	COMBINED    ListLtsConfigsRequestInstanceMode
}

func GetListLtsConfigsRequestInstanceModeEnum() ListLtsConfigsRequestInstanceModeEnum {
	return ListLtsConfigsRequestInstanceModeEnum{
		HA: ListLtsConfigsRequestInstanceMode{
			value: "Ha",
		},
		INDEPENDENT: ListLtsConfigsRequestInstanceMode{
			value: "Independent",
		},
		COMBINED: ListLtsConfigsRequestInstanceMode{
			value: "Combined",
		},
	}
}

func (c ListLtsConfigsRequestInstanceMode) Value() string {
	return c.value
}

func (c ListLtsConfigsRequestInstanceMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListLtsConfigsRequestInstanceMode) UnmarshalJSON(b []byte) error {
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
