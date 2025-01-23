package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListRiskTrendRequest Request Object
type ListRiskTrendRequest struct {

	// 语言
	XLanguage *ListRiskTrendRequestXLanguage `json:"X-Language,omitempty"`

	// 数据库类型
	DatastoreType string `json:"datastore_type"`

	// 开始时间
	StartAt int64 `json:"start_at"`

	// 结束时间
	EndAt int64 `json:"end_at"`

	// 指标码
	MetricCode string `json:"metric_code"`
}

func (o ListRiskTrendRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRiskTrendRequest struct{}"
	}

	return strings.Join([]string{"ListRiskTrendRequest", string(data)}, " ")
}

type ListRiskTrendRequestXLanguage struct {
	value string
}

type ListRiskTrendRequestXLanguageEnum struct {
	ZH_CN ListRiskTrendRequestXLanguage
	EN_US ListRiskTrendRequestXLanguage
}

func GetListRiskTrendRequestXLanguageEnum() ListRiskTrendRequestXLanguageEnum {
	return ListRiskTrendRequestXLanguageEnum{
		ZH_CN: ListRiskTrendRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListRiskTrendRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListRiskTrendRequestXLanguage) Value() string {
	return c.value
}

func (c ListRiskTrendRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRiskTrendRequestXLanguage) UnmarshalJSON(b []byte) error {
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
