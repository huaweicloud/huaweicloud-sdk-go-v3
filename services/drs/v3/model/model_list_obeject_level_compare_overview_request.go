package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListObejectLevelCompareOverviewRequest Request Object
type ListObejectLevelCompareOverviewRequest struct {

	// 请求语言类型。
	XLanguage *ListObejectLevelCompareOverviewRequestXLanguage `json:"X-Language,omitempty"`

	// 任务ID。
	JobId string `json:"job_id"`
}

func (o ListObejectLevelCompareOverviewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListObejectLevelCompareOverviewRequest struct{}"
	}

	return strings.Join([]string{"ListObejectLevelCompareOverviewRequest", string(data)}, " ")
}

type ListObejectLevelCompareOverviewRequestXLanguage struct {
	value string
}

type ListObejectLevelCompareOverviewRequestXLanguageEnum struct {
	EN_US ListObejectLevelCompareOverviewRequestXLanguage
	ZH_CN ListObejectLevelCompareOverviewRequestXLanguage
}

func GetListObejectLevelCompareOverviewRequestXLanguageEnum() ListObejectLevelCompareOverviewRequestXLanguageEnum {
	return ListObejectLevelCompareOverviewRequestXLanguageEnum{
		EN_US: ListObejectLevelCompareOverviewRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListObejectLevelCompareOverviewRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListObejectLevelCompareOverviewRequestXLanguage) Value() string {
	return c.value
}

func (c ListObejectLevelCompareOverviewRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListObejectLevelCompareOverviewRequestXLanguage) UnmarshalJSON(b []byte) error {
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
