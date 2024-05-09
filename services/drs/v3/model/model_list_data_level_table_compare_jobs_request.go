package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListDataLevelTableCompareJobsRequest Request Object
type ListDataLevelTableCompareJobsRequest struct {

	// 请求语言类型。
	XLanguage *ListDataLevelTableCompareJobsRequestXLanguage `json:"X-Language,omitempty"`

	// 任务ID。
	JobId string `json:"job_id"`

	// 每页显示的条目数量，最大值1000。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListDataLevelTableCompareJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataLevelTableCompareJobsRequest struct{}"
	}

	return strings.Join([]string{"ListDataLevelTableCompareJobsRequest", string(data)}, " ")
}

type ListDataLevelTableCompareJobsRequestXLanguage struct {
	value string
}

type ListDataLevelTableCompareJobsRequestXLanguageEnum struct {
	EN_US ListDataLevelTableCompareJobsRequestXLanguage
	ZH_CN ListDataLevelTableCompareJobsRequestXLanguage
}

func GetListDataLevelTableCompareJobsRequestXLanguageEnum() ListDataLevelTableCompareJobsRequestXLanguageEnum {
	return ListDataLevelTableCompareJobsRequestXLanguageEnum{
		EN_US: ListDataLevelTableCompareJobsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListDataLevelTableCompareJobsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListDataLevelTableCompareJobsRequestXLanguage) Value() string {
	return c.value
}

func (c ListDataLevelTableCompareJobsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDataLevelTableCompareJobsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
