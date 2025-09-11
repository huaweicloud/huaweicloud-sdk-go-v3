package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAppliedHistoriesRequest Request Object
type ListAppliedHistoriesRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**:   - zh-cn   - en-us  **默认取值**: en-us
	XLanguage *ListAppliedHistoriesRequestXLanguage `json:"X-Language,omitempty"`

	// 参数配置模板ID。
	ConfigId string `json:"config_id"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。默认为100，不能为负数，最小值为1，最大值为100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAppliedHistoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppliedHistoriesRequest struct{}"
	}

	return strings.Join([]string{"ListAppliedHistoriesRequest", string(data)}, " ")
}

type ListAppliedHistoriesRequestXLanguage struct {
	value string
}

type ListAppliedHistoriesRequestXLanguageEnum struct {
	ZH_CN ListAppliedHistoriesRequestXLanguage
	EN_US ListAppliedHistoriesRequestXLanguage
}

func GetListAppliedHistoriesRequestXLanguageEnum() ListAppliedHistoriesRequestXLanguageEnum {
	return ListAppliedHistoriesRequestXLanguageEnum{
		ZH_CN: ListAppliedHistoriesRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListAppliedHistoriesRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListAppliedHistoriesRequestXLanguage) Value() string {
	return c.value
}

func (c ListAppliedHistoriesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAppliedHistoriesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
