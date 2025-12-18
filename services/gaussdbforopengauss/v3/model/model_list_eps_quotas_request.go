package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListEpsQuotasRequest Request Object
type ListEpsQuotasRequest struct {

	// 语言, 默认值为en-us。
	XLanguage *ListEpsQuotasRequestXLanguage `json:"X-Language,omitempty"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。默认为100，不能为负数，最小值为1，最大值为100。
	Limit *int32 `json:"limit,omitempty"`

	// 企业项目ID。 - 对于未开通企业多项目服务的用户，不传该参数。 - 对于已开通企业多项目服务的用户，不传该参数时，表示为default企业项目。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 企业项目名称。
	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`
}

func (o ListEpsQuotasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEpsQuotasRequest struct{}"
	}

	return strings.Join([]string{"ListEpsQuotasRequest", string(data)}, " ")
}

type ListEpsQuotasRequestXLanguage struct {
	value string
}

type ListEpsQuotasRequestXLanguageEnum struct {
	ZH_CN ListEpsQuotasRequestXLanguage
	EN_US ListEpsQuotasRequestXLanguage
}

func GetListEpsQuotasRequestXLanguageEnum() ListEpsQuotasRequestXLanguageEnum {
	return ListEpsQuotasRequestXLanguageEnum{
		ZH_CN: ListEpsQuotasRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListEpsQuotasRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListEpsQuotasRequestXLanguage) Value() string {
	return c.value
}

func (c ListEpsQuotasRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEpsQuotasRequestXLanguage) UnmarshalJSON(b []byte) error {
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
