package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowEnterpriseProjectRequest Request Object
type ShowEnterpriseProjectRequest struct {

	// 请求语言类型。
	XLanguage *ShowEnterpriseProjectRequestXLanguage `json:"X-Language,omitempty"`

	// 偏移量，表示从此偏移量开始查询， offset 大于等于 0。默认为0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。默认为10，取值范围【1-1000】
	Limit *int32 `json:"limit,omitempty"`

	// 企业项目名称，支持模糊搜索。
	Name *string `json:"name,omitempty"`

	// IAM用户所属帐号ID。op_service权限必须携带此参数，非op_service权限可不携带此参数。
	DomainId *string `json:"domain_id,omitempty"`
}

func (o ShowEnterpriseProjectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEnterpriseProjectRequest struct{}"
	}

	return strings.Join([]string{"ShowEnterpriseProjectRequest", string(data)}, " ")
}

type ShowEnterpriseProjectRequestXLanguage struct {
	value string
}

type ShowEnterpriseProjectRequestXLanguageEnum struct {
	EN_US ShowEnterpriseProjectRequestXLanguage
	ZH_CN ShowEnterpriseProjectRequestXLanguage
}

func GetShowEnterpriseProjectRequestXLanguageEnum() ShowEnterpriseProjectRequestXLanguageEnum {
	return ShowEnterpriseProjectRequestXLanguageEnum{
		EN_US: ShowEnterpriseProjectRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowEnterpriseProjectRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowEnterpriseProjectRequestXLanguage) Value() string {
	return c.value
}

func (c ShowEnterpriseProjectRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowEnterpriseProjectRequestXLanguage) UnmarshalJSON(b []byte) error {
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
