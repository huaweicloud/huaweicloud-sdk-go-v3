package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowFilesRequest Request Object
type ShowFilesRequest struct {

	// 请求语言类型。
	XLanguage *ShowFilesRequestXLanguage `json:"X-Language,omitempty"`

	// 偏移量，表示查询该偏移量后面的记录。
	Offset *int32 `json:"offset,omitempty"`

	// 查询返回记录的数量限制。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowFilesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFilesRequest struct{}"
	}

	return strings.Join([]string{"ShowFilesRequest", string(data)}, " ")
}

type ShowFilesRequestXLanguage struct {
	value string
}

type ShowFilesRequestXLanguageEnum struct {
	EN_US ShowFilesRequestXLanguage
	ZH_CN ShowFilesRequestXLanguage
	ES_US ShowFilesRequestXLanguage
	PT_BR ShowFilesRequestXLanguage
}

func GetShowFilesRequestXLanguageEnum() ShowFilesRequestXLanguageEnum {
	return ShowFilesRequestXLanguageEnum{
		EN_US: ShowFilesRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowFilesRequestXLanguage{
			value: "zh-cn",
		},
		ES_US: ShowFilesRequestXLanguage{
			value: "es-us",
		},
		PT_BR: ShowFilesRequestXLanguage{
			value: "pt-br",
		},
	}
}

func (c ShowFilesRequestXLanguage) Value() string {
	return c.value
}

func (c ShowFilesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowFilesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
