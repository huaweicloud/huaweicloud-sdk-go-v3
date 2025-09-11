package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListInstanceEngineDetailRequest Request Object
type ListInstanceEngineDetailRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**:   - zh-cn   - en-us  **默认取值**: en-us
	XLanguage *ListInstanceEngineDetailRequestXLanguage `json:"X-Language,omitempty"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。默认为100，不能为负数，最小值为1，最大值为100
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListInstanceEngineDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceEngineDetailRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceEngineDetailRequest", string(data)}, " ")
}

type ListInstanceEngineDetailRequestXLanguage struct {
	value string
}

type ListInstanceEngineDetailRequestXLanguageEnum struct {
	ZH_CN ListInstanceEngineDetailRequestXLanguage
	EN_US ListInstanceEngineDetailRequestXLanguage
}

func GetListInstanceEngineDetailRequestXLanguageEnum() ListInstanceEngineDetailRequestXLanguageEnum {
	return ListInstanceEngineDetailRequestXLanguageEnum{
		ZH_CN: ListInstanceEngineDetailRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListInstanceEngineDetailRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListInstanceEngineDetailRequestXLanguage) Value() string {
	return c.value
}

func (c ListInstanceEngineDetailRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstanceEngineDetailRequestXLanguage) UnmarshalJSON(b []byte) error {
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
