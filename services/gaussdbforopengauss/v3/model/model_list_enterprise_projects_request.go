package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListEnterpriseProjectsRequest Request Object
type ListEnterpriseProjectsRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *ListEnterpriseProjectsRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 企业项目名称关键字。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	NameKeyword *string `json:"name_keyword,omitempty"`

	// **参数解释**: 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询。 **约束限制**: 必须为数字，不能为负数。 **取值范围**: 不涉及。 **默认取值**: 默认为0（偏移0条数据，表示从第一条数据开始查询）。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 查询记录数。 **约束限制**: 不能为负数。 **取值范围**: 最小值为1，最大值为1000。 **默认取值**: 100
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListEnterpriseProjectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnterpriseProjectsRequest struct{}"
	}

	return strings.Join([]string{"ListEnterpriseProjectsRequest", string(data)}, " ")
}

type ListEnterpriseProjectsRequestXLanguage struct {
	value string
}

type ListEnterpriseProjectsRequestXLanguageEnum struct {
	ZH_CN ListEnterpriseProjectsRequestXLanguage
	EN_US ListEnterpriseProjectsRequestXLanguage
}

func GetListEnterpriseProjectsRequestXLanguageEnum() ListEnterpriseProjectsRequestXLanguageEnum {
	return ListEnterpriseProjectsRequestXLanguageEnum{
		ZH_CN: ListEnterpriseProjectsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListEnterpriseProjectsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListEnterpriseProjectsRequestXLanguage) Value() string {
	return c.value
}

func (c ListEnterpriseProjectsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEnterpriseProjectsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
