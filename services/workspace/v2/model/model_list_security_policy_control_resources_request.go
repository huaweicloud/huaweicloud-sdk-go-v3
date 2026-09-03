package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSecurityPolicyControlResourcesRequest Request Object
type ListSecurityPolicyControlResourcesRequest struct {

	// 语言，用于国际化。 - en-us：英文 - zh-cn：中文
	XLanguage *ListSecurityPolicyControlResourcesRequestXLanguage `json:"X-Language,omitempty"`

	// 偏移量，默认0。
	Offset *int32 `json:"offset,omitempty"`

	// 分页大小，默认20。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListSecurityPolicyControlResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityPolicyControlResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListSecurityPolicyControlResourcesRequest", string(data)}, " ")
}

type ListSecurityPolicyControlResourcesRequestXLanguage struct {
	value string
}

type ListSecurityPolicyControlResourcesRequestXLanguageEnum struct {
	EN_US ListSecurityPolicyControlResourcesRequestXLanguage
	ZH_CN ListSecurityPolicyControlResourcesRequestXLanguage
}

func GetListSecurityPolicyControlResourcesRequestXLanguageEnum() ListSecurityPolicyControlResourcesRequestXLanguageEnum {
	return ListSecurityPolicyControlResourcesRequestXLanguageEnum{
		EN_US: ListSecurityPolicyControlResourcesRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListSecurityPolicyControlResourcesRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListSecurityPolicyControlResourcesRequestXLanguage) Value() string {
	return c.value
}

func (c ListSecurityPolicyControlResourcesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSecurityPolicyControlResourcesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
