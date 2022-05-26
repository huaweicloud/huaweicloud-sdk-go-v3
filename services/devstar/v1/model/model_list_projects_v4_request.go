package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListProjectsV4Request struct {

	// 语言类型 中文:zh-cn 英文:en-us
	XLanguage *ListProjectsV4RequestXLanguage `json:"X-Language,omitempty"`

	// 搜索关键字,支持按名称和描述搜索，默认null
	Keyword *string `json:"keyword,omitempty"`

	// 每页显示的条目数量,默认100
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询,默认0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListProjectsV4Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectsV4Request struct{}"
	}

	return strings.Join([]string{"ListProjectsV4Request", string(data)}, " ")
}

type ListProjectsV4RequestXLanguage struct {
	value string
}

type ListProjectsV4RequestXLanguageEnum struct {
	ZH_CN ListProjectsV4RequestXLanguage
	EN_US ListProjectsV4RequestXLanguage
}

func GetListProjectsV4RequestXLanguageEnum() ListProjectsV4RequestXLanguageEnum {
	return ListProjectsV4RequestXLanguageEnum{
		ZH_CN: ListProjectsV4RequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListProjectsV4RequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListProjectsV4RequestXLanguage) Value() string {
	return c.value
}

func (c ListProjectsV4RequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListProjectsV4RequestXLanguage) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
