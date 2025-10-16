package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSessionTopSqlStatisticsRequest Request Object
type ListSessionTopSqlStatisticsRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *ListSessionTopSqlStatisticsRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**: 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询。例如：该参数指定为0，limit指定为10，则只展示第1~10条数据。 **约束限制**: 不涉及。 **取值范围**: [0, 2^31-1] **默认取值**: 默认为0（偏移0条数据，表示从第一条数据开始查询）。
	Offset int32 `json:"offset"`

	// **参数解释**: 查询记录数。例如该参数设定为10，则查询结果最多只显示10条记录。 **约束限制**: 不涉及。 **取值范围**: [1, 100] **默认取值**: 默认为100。
	Limit int32 `json:"limit"`
}

func (o ListSessionTopSqlStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSessionTopSqlStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ListSessionTopSqlStatisticsRequest", string(data)}, " ")
}

type ListSessionTopSqlStatisticsRequestXLanguage struct {
	value string
}

type ListSessionTopSqlStatisticsRequestXLanguageEnum struct {
	ZH_CN ListSessionTopSqlStatisticsRequestXLanguage
	EN_US ListSessionTopSqlStatisticsRequestXLanguage
}

func GetListSessionTopSqlStatisticsRequestXLanguageEnum() ListSessionTopSqlStatisticsRequestXLanguageEnum {
	return ListSessionTopSqlStatisticsRequestXLanguageEnum{
		ZH_CN: ListSessionTopSqlStatisticsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListSessionTopSqlStatisticsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListSessionTopSqlStatisticsRequestXLanguage) Value() string {
	return c.value
}

func (c ListSessionTopSqlStatisticsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSessionTopSqlStatisticsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
