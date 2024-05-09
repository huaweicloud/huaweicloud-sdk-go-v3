package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListContentCompareDifferenceRequest Request Object
type ListContentCompareDifferenceRequest struct {

	// 请求语言类型。
	XLanguage *ListContentCompareDifferenceRequestXLanguage `json:"X-Language,omitempty"`

	// 任务ID。
	JobId string `json:"job_id"`

	// 对比任务ID。
	CompareJobId string `json:"compare_job_id"`

	// 表名。
	TableName *string `json:"table_name,omitempty"`

	// 源库名称。
	DbName *string `json:"db_name,omitempty"`

	// 每页显示的条目数量，最大值1000。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListContentCompareDifferenceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListContentCompareDifferenceRequest struct{}"
	}

	return strings.Join([]string{"ListContentCompareDifferenceRequest", string(data)}, " ")
}

type ListContentCompareDifferenceRequestXLanguage struct {
	value string
}

type ListContentCompareDifferenceRequestXLanguageEnum struct {
	EN_US ListContentCompareDifferenceRequestXLanguage
	ZH_CN ListContentCompareDifferenceRequestXLanguage
}

func GetListContentCompareDifferenceRequestXLanguageEnum() ListContentCompareDifferenceRequestXLanguageEnum {
	return ListContentCompareDifferenceRequestXLanguageEnum{
		EN_US: ListContentCompareDifferenceRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListContentCompareDifferenceRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListContentCompareDifferenceRequestXLanguage) Value() string {
	return c.value
}

func (c ListContentCompareDifferenceRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListContentCompareDifferenceRequestXLanguage) UnmarshalJSON(b []byte) error {
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
