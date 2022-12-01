package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ShowDbObjectTemplateProgressRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *ShowDbObjectTemplateProgressRequestXLanguage `json:"X-Language,omitempty"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowDbObjectTemplateProgressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDbObjectTemplateProgressRequest struct{}"
	}

	return strings.Join([]string{"ShowDbObjectTemplateProgressRequest", string(data)}, " ")
}

type ShowDbObjectTemplateProgressRequestXLanguage struct {
	value string
}

type ShowDbObjectTemplateProgressRequestXLanguageEnum struct {
	EN_US ShowDbObjectTemplateProgressRequestXLanguage
	ZH_CN ShowDbObjectTemplateProgressRequestXLanguage
}

func GetShowDbObjectTemplateProgressRequestXLanguageEnum() ShowDbObjectTemplateProgressRequestXLanguageEnum {
	return ShowDbObjectTemplateProgressRequestXLanguageEnum{
		EN_US: ShowDbObjectTemplateProgressRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowDbObjectTemplateProgressRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowDbObjectTemplateProgressRequestXLanguage) Value() string {
	return c.value
}

func (c ShowDbObjectTemplateProgressRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDbObjectTemplateProgressRequestXLanguage) UnmarshalJSON(b []byte) error {
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
