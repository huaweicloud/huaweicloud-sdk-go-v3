package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListContentCompareDetailRequest Request Object
type ListContentCompareDetailRequest struct {

	// 请求语言类型。
	XLanguage *ListContentCompareDetailRequestXLanguage `json:"X-Language,omitempty"`

	// 任务ID。
	JobId string `json:"job_id"`

	// 对比任务ID。
	CompareJobId string `json:"compare_job_id"`

	// 目标库名称。
	TargetDbName *string `json:"target_db_name,omitempty"`

	// 源库名称。
	DbName *string `json:"db_name,omitempty"`

	// 类型。 compare：对比 unCompare：无法对比
	Type *string `json:"type,omitempty"`

	// 每页显示的条目数量，最大值1000。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListContentCompareDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListContentCompareDetailRequest struct{}"
	}

	return strings.Join([]string{"ListContentCompareDetailRequest", string(data)}, " ")
}

type ListContentCompareDetailRequestXLanguage struct {
	value string
}

type ListContentCompareDetailRequestXLanguageEnum struct {
	EN_US ListContentCompareDetailRequestXLanguage
	ZH_CN ListContentCompareDetailRequestXLanguage
}

func GetListContentCompareDetailRequestXLanguageEnum() ListContentCompareDetailRequestXLanguageEnum {
	return ListContentCompareDetailRequestXLanguageEnum{
		EN_US: ListContentCompareDetailRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListContentCompareDetailRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListContentCompareDetailRequestXLanguage) Value() string {
	return c.value
}

func (c ListContentCompareDetailRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListContentCompareDetailRequestXLanguage) UnmarshalJSON(b []byte) error {
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
