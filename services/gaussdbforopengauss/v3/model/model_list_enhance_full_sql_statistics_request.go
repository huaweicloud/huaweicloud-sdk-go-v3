package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListEnhanceFullSqlStatisticsRequest Request Object
type ListEnhanceFullSqlStatisticsRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *ListEnhanceFullSqlStatisticsRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	Body *ListEnhanceFullSqlStatisticsRequestBody `json:"body,omitempty"`
}

func (o ListEnhanceFullSqlStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnhanceFullSqlStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ListEnhanceFullSqlStatisticsRequest", string(data)}, " ")
}

type ListEnhanceFullSqlStatisticsRequestXLanguage struct {
	value string
}

type ListEnhanceFullSqlStatisticsRequestXLanguageEnum struct {
	ZH_CN ListEnhanceFullSqlStatisticsRequestXLanguage
	EN_US ListEnhanceFullSqlStatisticsRequestXLanguage
}

func GetListEnhanceFullSqlStatisticsRequestXLanguageEnum() ListEnhanceFullSqlStatisticsRequestXLanguageEnum {
	return ListEnhanceFullSqlStatisticsRequestXLanguageEnum{
		ZH_CN: ListEnhanceFullSqlStatisticsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListEnhanceFullSqlStatisticsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListEnhanceFullSqlStatisticsRequestXLanguage) Value() string {
	return c.value
}

func (c ListEnhanceFullSqlStatisticsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEnhanceFullSqlStatisticsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
