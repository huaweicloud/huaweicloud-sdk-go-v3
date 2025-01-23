package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListInstanceMultiNodesSingleMetricRequest Request Object
type ListInstanceMultiNodesSingleMetricRequest struct {

	// 语言
	XLanguage *ListInstanceMultiNodesSingleMetricRequestXLanguage `json:"X-Language,omitempty"`

	Body *ListInstanceMultiNodesSingleMetric `json:"body,omitempty"`
}

func (o ListInstanceMultiNodesSingleMetricRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceMultiNodesSingleMetricRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceMultiNodesSingleMetricRequest", string(data)}, " ")
}

type ListInstanceMultiNodesSingleMetricRequestXLanguage struct {
	value string
}

type ListInstanceMultiNodesSingleMetricRequestXLanguageEnum struct {
	ZH_CN ListInstanceMultiNodesSingleMetricRequestXLanguage
	EN_US ListInstanceMultiNodesSingleMetricRequestXLanguage
}

func GetListInstanceMultiNodesSingleMetricRequestXLanguageEnum() ListInstanceMultiNodesSingleMetricRequestXLanguageEnum {
	return ListInstanceMultiNodesSingleMetricRequestXLanguageEnum{
		ZH_CN: ListInstanceMultiNodesSingleMetricRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListInstanceMultiNodesSingleMetricRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListInstanceMultiNodesSingleMetricRequestXLanguage) Value() string {
	return c.value
}

func (c ListInstanceMultiNodesSingleMetricRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstanceMultiNodesSingleMetricRequestXLanguage) UnmarshalJSON(b []byte) error {
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
