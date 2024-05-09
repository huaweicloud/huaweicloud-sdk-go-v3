package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListContentCompareOverviewRequest Request Object
type ListContentCompareOverviewRequest struct {

	// 请求语言类型。
	XLanguage *ListContentCompareOverviewRequestXLanguage `json:"X-Language,omitempty"`

	// 任务ID。
	JobId string `json:"job_id"`

	// 对比任务ID。
	CompareJobId string `json:"compare_job_id"`

	// 每页显示的条目数量，最大值1000。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListContentCompareOverviewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListContentCompareOverviewRequest struct{}"
	}

	return strings.Join([]string{"ListContentCompareOverviewRequest", string(data)}, " ")
}

type ListContentCompareOverviewRequestXLanguage struct {
	value string
}

type ListContentCompareOverviewRequestXLanguageEnum struct {
	EN_US ListContentCompareOverviewRequestXLanguage
	ZH_CN ListContentCompareOverviewRequestXLanguage
}

func GetListContentCompareOverviewRequestXLanguageEnum() ListContentCompareOverviewRequestXLanguageEnum {
	return ListContentCompareOverviewRequestXLanguageEnum{
		EN_US: ListContentCompareOverviewRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListContentCompareOverviewRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListContentCompareOverviewRequestXLanguage) Value() string {
	return c.value
}

func (c ListContentCompareOverviewRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListContentCompareOverviewRequestXLanguage) UnmarshalJSON(b []byte) error {
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
