package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListDatabaseVersionsRequest Request Object
type ListDatabaseVersionsRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *ListDatabaseVersionsRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询。例如：该参数指定为1，limit指定为10，则只展示第2-11条数据。 **约束限制**: 不涉及。 **取值范围**: [0, 2^31-1] **默认取值**: 不涉及。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 查询记录数。例如该参数设定为10，则查询结果最多只显示10条记录。 **约束限制**: 不涉及。 **取值范围**: [1, 100] **默认取值**: 100
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListDatabaseVersionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabaseVersionsRequest struct{}"
	}

	return strings.Join([]string{"ListDatabaseVersionsRequest", string(data)}, " ")
}

type ListDatabaseVersionsRequestXLanguage struct {
	value string
}

type ListDatabaseVersionsRequestXLanguageEnum struct {
	ZH_CN ListDatabaseVersionsRequestXLanguage
	EN_US ListDatabaseVersionsRequestXLanguage
}

func GetListDatabaseVersionsRequestXLanguageEnum() ListDatabaseVersionsRequestXLanguageEnum {
	return ListDatabaseVersionsRequestXLanguageEnum{
		ZH_CN: ListDatabaseVersionsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListDatabaseVersionsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListDatabaseVersionsRequestXLanguage) Value() string {
	return c.value
}

func (c ListDatabaseVersionsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDatabaseVersionsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
