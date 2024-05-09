package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowHealthObjectCompareJobOverviewRequest Request Object
type ShowHealthObjectCompareJobOverviewRequest struct {

	// 请求语言类型。
	XLanguage *ShowHealthObjectCompareJobOverviewRequestXLanguage `json:"X-Language,omitempty"`

	// 任务ID。
	JobId string `json:"job_id"`

	// 对比任务ID。
	CompareJobId string `json:"compare_job_id"`
}

func (o ShowHealthObjectCompareJobOverviewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHealthObjectCompareJobOverviewRequest struct{}"
	}

	return strings.Join([]string{"ShowHealthObjectCompareJobOverviewRequest", string(data)}, " ")
}

type ShowHealthObjectCompareJobOverviewRequestXLanguage struct {
	value string
}

type ShowHealthObjectCompareJobOverviewRequestXLanguageEnum struct {
	EN_US ShowHealthObjectCompareJobOverviewRequestXLanguage
	ZH_CN ShowHealthObjectCompareJobOverviewRequestXLanguage
}

func GetShowHealthObjectCompareJobOverviewRequestXLanguageEnum() ShowHealthObjectCompareJobOverviewRequestXLanguageEnum {
	return ShowHealthObjectCompareJobOverviewRequestXLanguageEnum{
		EN_US: ShowHealthObjectCompareJobOverviewRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowHealthObjectCompareJobOverviewRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowHealthObjectCompareJobOverviewRequestXLanguage) Value() string {
	return c.value
}

func (c ShowHealthObjectCompareJobOverviewRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowHealthObjectCompareJobOverviewRequestXLanguage) UnmarshalJSON(b []byte) error {
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
