package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowHealthCompareJobDetailRequest Request Object
type ShowHealthCompareJobDetailRequest struct {

	// 请求语言类型。
	XLanguage *ShowHealthCompareJobDetailRequestXLanguage `json:"X-Language,omitempty"`

	// 任务ID。
	JobId string `json:"job_id"`

	// 对比任务ID。
	CompareJobId string `json:"compare_job_id"`
}

func (o ShowHealthCompareJobDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHealthCompareJobDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowHealthCompareJobDetailRequest", string(data)}, " ")
}

type ShowHealthCompareJobDetailRequestXLanguage struct {
	value string
}

type ShowHealthCompareJobDetailRequestXLanguageEnum struct {
	EN_US ShowHealthCompareJobDetailRequestXLanguage
	ZH_CN ShowHealthCompareJobDetailRequestXLanguage
}

func GetShowHealthCompareJobDetailRequestXLanguageEnum() ShowHealthCompareJobDetailRequestXLanguageEnum {
	return ShowHealthCompareJobDetailRequestXLanguageEnum{
		EN_US: ShowHealthCompareJobDetailRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowHealthCompareJobDetailRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowHealthCompareJobDetailRequestXLanguage) Value() string {
	return c.value
}

func (c ShowHealthCompareJobDetailRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowHealthCompareJobDetailRequestXLanguage) UnmarshalJSON(b []byte) error {
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
