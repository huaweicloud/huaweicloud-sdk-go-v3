package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListTemplateViewHistoriesRequest Request Object
type ListTemplateViewHistoriesRequest struct {

	// 语言类型，缺省值为“zh-cn”。  枚举值： - zh-cn：中文 - en-us：英文
	XLanguage *ListTemplateViewHistoriesRequestXLanguage `json:"X-Language,omitempty"`

	// 平台来源： - 0：查询CodeLabs中用户浏览过的模板。 - 1：查询DevStar中用户浏览过的模板。
	PlatformSource ListTemplateViewHistoriesRequestPlatformSource `json:"platform_source"`
}

func (o ListTemplateViewHistoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTemplateViewHistoriesRequest struct{}"
	}

	return strings.Join([]string{"ListTemplateViewHistoriesRequest", string(data)}, " ")
}

type ListTemplateViewHistoriesRequestXLanguage struct {
	value string
}

type ListTemplateViewHistoriesRequestXLanguageEnum struct {
	ZH_CN ListTemplateViewHistoriesRequestXLanguage
	EN_US ListTemplateViewHistoriesRequestXLanguage
}

func GetListTemplateViewHistoriesRequestXLanguageEnum() ListTemplateViewHistoriesRequestXLanguageEnum {
	return ListTemplateViewHistoriesRequestXLanguageEnum{
		ZH_CN: ListTemplateViewHistoriesRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListTemplateViewHistoriesRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListTemplateViewHistoriesRequestXLanguage) Value() string {
	return c.value
}

func (c ListTemplateViewHistoriesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTemplateViewHistoriesRequestXLanguage) UnmarshalJSON(b []byte) error {
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

type ListTemplateViewHistoriesRequestPlatformSource struct {
	value int32
}

type ListTemplateViewHistoriesRequestPlatformSourceEnum struct {
	E_0 ListTemplateViewHistoriesRequestPlatformSource
	E_1 ListTemplateViewHistoriesRequestPlatformSource
}

func GetListTemplateViewHistoriesRequestPlatformSourceEnum() ListTemplateViewHistoriesRequestPlatformSourceEnum {
	return ListTemplateViewHistoriesRequestPlatformSourceEnum{
		E_0: ListTemplateViewHistoriesRequestPlatformSource{
			value: 0,
		}, E_1: ListTemplateViewHistoriesRequestPlatformSource{
			value: 1,
		},
	}
}

func (c ListTemplateViewHistoriesRequestPlatformSource) Value() int32 {
	return c.value
}

func (c ListTemplateViewHistoriesRequestPlatformSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTemplateViewHistoriesRequestPlatformSource) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
