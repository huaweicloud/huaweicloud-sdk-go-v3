package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowCompareProgressRequest Request Object
type ShowCompareProgressRequest struct {

	// 请求语言类型。 -en-us：英文 -zh-cn：中文
	XLanguage *ShowCompareProgressRequestXLanguage `json:"X-Language,omitempty"`

	// 任务ID。
	JobId string `json:"job_id"`

	// 对比任务ID。
	CompareJobId string `json:"compare_job_id"`
}

func (o ShowCompareProgressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCompareProgressRequest struct{}"
	}

	return strings.Join([]string{"ShowCompareProgressRequest", string(data)}, " ")
}

type ShowCompareProgressRequestXLanguage struct {
	value string
}

type ShowCompareProgressRequestXLanguageEnum struct {
	EN_US ShowCompareProgressRequestXLanguage
	ZH_CN ShowCompareProgressRequestXLanguage
}

func GetShowCompareProgressRequestXLanguageEnum() ShowCompareProgressRequestXLanguageEnum {
	return ShowCompareProgressRequestXLanguageEnum{
		EN_US: ShowCompareProgressRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowCompareProgressRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowCompareProgressRequestXLanguage) Value() string {
	return c.value
}

func (c ShowCompareProgressRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowCompareProgressRequestXLanguage) UnmarshalJSON(b []byte) error {
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
