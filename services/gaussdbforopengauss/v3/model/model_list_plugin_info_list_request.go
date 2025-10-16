package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListPluginInfoListRequest Request Object
type ListPluginInfoListRequest struct {

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**: 用户Token。 通过调用IAM服务[获取用户token](https://support.huaweicloud.com/intl/zh-cn/api-iam/iam_30_0001.html)。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	XLanguage *ListPluginInfoListRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询。例如：该参数指定为1，limit指定为10，则只展示第2-11条数据。 **约束限制**: 不涉及。 **取值范围**: [0, 2^31-1] **默认取值**: 默认为0（偏移0条数据，表示从第一条数据开始查询）。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 查询记录数。例如该参数设定为10，则查询结果最多只显示10条记录。 **约束限制**: 不涉及。 **取值范围**: [1, 100] **默认取值**: 默认为100。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 插件包名称。 **约束限制**: 不涉及。 **取值范围**: - postgis  **默认取值**: 不涉及。
	PluginName *ListPluginInfoListRequestPluginName `json:"plugin_name,omitempty"`
}

func (o ListPluginInfoListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPluginInfoListRequest struct{}"
	}

	return strings.Join([]string{"ListPluginInfoListRequest", string(data)}, " ")
}

type ListPluginInfoListRequestXLanguage struct {
	value string
}

type ListPluginInfoListRequestXLanguageEnum struct {
	ZH_CN ListPluginInfoListRequestXLanguage
	EN_US ListPluginInfoListRequestXLanguage
}

func GetListPluginInfoListRequestXLanguageEnum() ListPluginInfoListRequestXLanguageEnum {
	return ListPluginInfoListRequestXLanguageEnum{
		ZH_CN: ListPluginInfoListRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListPluginInfoListRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListPluginInfoListRequestXLanguage) Value() string {
	return c.value
}

func (c ListPluginInfoListRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPluginInfoListRequestXLanguage) UnmarshalJSON(b []byte) error {
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

type ListPluginInfoListRequestPluginName struct {
	value string
}

type ListPluginInfoListRequestPluginNameEnum struct {
	POSTGIS ListPluginInfoListRequestPluginName
}

func GetListPluginInfoListRequestPluginNameEnum() ListPluginInfoListRequestPluginNameEnum {
	return ListPluginInfoListRequestPluginNameEnum{
		POSTGIS: ListPluginInfoListRequestPluginName{
			value: "postgis",
		},
	}
}

func (c ListPluginInfoListRequestPluginName) Value() string {
	return c.value
}

func (c ListPluginInfoListRequestPluginName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPluginInfoListRequestPluginName) UnmarshalJSON(b []byte) error {
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
