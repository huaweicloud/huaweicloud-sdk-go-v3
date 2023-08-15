package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowProgressDataRequest Request Object
type ShowProgressDataRequest struct {

	// 请求语言类型。
	XLanguage *ShowProgressDataRequestXLanguage `json:"X-Language,omitempty"`

	// 任务ID。
	JobId string `json:"job_id"`

	// 偏移量，表示从此偏移量开始查询， offset 大于等于 0。默认为0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。默认为10，取值范围【1-1000】
	Limit *int32 `json:"limit,omitempty"`

	// 迁移对象类型。 - table - event - table_structure - procedure - view - function - database - trigger - table_indexs
	Type string `json:"type"`
}

func (o ShowProgressDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProgressDataRequest struct{}"
	}

	return strings.Join([]string{"ShowProgressDataRequest", string(data)}, " ")
}

type ShowProgressDataRequestXLanguage struct {
	value string
}

type ShowProgressDataRequestXLanguageEnum struct {
	EN_US ShowProgressDataRequestXLanguage
	ZH_CN ShowProgressDataRequestXLanguage
}

func GetShowProgressDataRequestXLanguageEnum() ShowProgressDataRequestXLanguageEnum {
	return ShowProgressDataRequestXLanguageEnum{
		EN_US: ShowProgressDataRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowProgressDataRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowProgressDataRequestXLanguage) Value() string {
	return c.value
}

func (c ShowProgressDataRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowProgressDataRequestXLanguage) UnmarshalJSON(b []byte) error {
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
