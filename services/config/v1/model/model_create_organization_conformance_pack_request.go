package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateOrganizationConformancePackRequest Request Object
type CreateOrganizationConformancePackRequest struct {

	// 组织ID。
	OrganizationId string `json:"organization_id"`

	// 组织合规包信息语言，默认为\"en-us\"英文
	XLanguage *CreateOrganizationConformancePackRequestXLanguage `json:"X-Language,omitempty"`

	Body *OrgConformancePackRequestBody `json:"body,omitempty"`
}

func (o CreateOrganizationConformancePackRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrganizationConformancePackRequest struct{}"
	}

	return strings.Join([]string{"CreateOrganizationConformancePackRequest", string(data)}, " ")
}

type CreateOrganizationConformancePackRequestXLanguage struct {
	value string
}

type CreateOrganizationConformancePackRequestXLanguageEnum struct {
	ZH_CN CreateOrganizationConformancePackRequestXLanguage
	EN_US CreateOrganizationConformancePackRequestXLanguage
}

func GetCreateOrganizationConformancePackRequestXLanguageEnum() CreateOrganizationConformancePackRequestXLanguageEnum {
	return CreateOrganizationConformancePackRequestXLanguageEnum{
		ZH_CN: CreateOrganizationConformancePackRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: CreateOrganizationConformancePackRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c CreateOrganizationConformancePackRequestXLanguage) Value() string {
	return c.value
}

func (c CreateOrganizationConformancePackRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateOrganizationConformancePackRequestXLanguage) UnmarshalJSON(b []byte) error {
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
