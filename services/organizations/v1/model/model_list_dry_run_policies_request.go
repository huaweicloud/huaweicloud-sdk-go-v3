package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListDryRunPoliciesRequest Request Object
type ListDryRunPoliciesRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 根、组织单元或账号的唯一标识符（ID）。
	AttachedEntityId *string `json:"attached_entity_id,omitempty"`

	// 页面中最大结果数量。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记。
	Marker *string `json:"marker,omitempty"`

	// 选择接口返回的信息的语言
	XLanguage *ListDryRunPoliciesRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ListDryRunPoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDryRunPoliciesRequest struct{}"
	}

	return strings.Join([]string{"ListDryRunPoliciesRequest", string(data)}, " ")
}

type ListDryRunPoliciesRequestXLanguage struct {
	value string
}

type ListDryRunPoliciesRequestXLanguageEnum struct {
	ZH_CN ListDryRunPoliciesRequestXLanguage
	EN_US ListDryRunPoliciesRequestXLanguage
}

func GetListDryRunPoliciesRequestXLanguageEnum() ListDryRunPoliciesRequestXLanguageEnum {
	return ListDryRunPoliciesRequestXLanguageEnum{
		ZH_CN: ListDryRunPoliciesRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListDryRunPoliciesRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListDryRunPoliciesRequestXLanguage) Value() string {
	return c.value
}

func (c ListDryRunPoliciesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDryRunPoliciesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
