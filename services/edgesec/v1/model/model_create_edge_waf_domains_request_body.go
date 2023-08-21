package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateEdgeWafDomainsRequestBody 创建防护域名的请求
type CreateEdgeWafDomainsRequestBody struct {

	// 防护域名（可带端口），通过查询CDN域名接口获取
	DomainName string `json:"domain_name"`

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 防护域名关联的策略id，通过查询WAF防护策略接口获取
	PolicyId *string `json:"policy_id,omitempty"`

	// 证书id，通过查询证书列表接口（ListCertificates）接口获取证书id   - 对外协议为HTTP时不需要填写   - 对外协议HTTPS时为必填参数   - 查询证书列表接口未开放时，从边缘安全控制台->边缘WAF->证书管理获取
	CertificateId *string `json:"certificate_id,omitempty"`

	// 域名名称
	WebTag *string `json:"web_tag,omitempty"`

	// 域名描述
	Description *string `json:"description,omitempty"`

	// 域名在CDN所属区域，通过查询CDN域名接口获取
	AreaType CreateEdgeWafDomainsRequestBodyAreaType `json:"area_type"`
}

func (o CreateEdgeWafDomainsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEdgeWafDomainsRequestBody struct{}"
	}

	return strings.Join([]string{"CreateEdgeWafDomainsRequestBody", string(data)}, " ")
}

type CreateEdgeWafDomainsRequestBodyAreaType struct {
	value string
}

type CreateEdgeWafDomainsRequestBodyAreaTypeEnum struct {
	MAINLAND_CHINA         CreateEdgeWafDomainsRequestBodyAreaType
	OUTSIDE_MAINLAND_CHINA CreateEdgeWafDomainsRequestBodyAreaType
	GLOBAL                 CreateEdgeWafDomainsRequestBodyAreaType
	EUROPE                 CreateEdgeWafDomainsRequestBodyAreaType
}

func GetCreateEdgeWafDomainsRequestBodyAreaTypeEnum() CreateEdgeWafDomainsRequestBodyAreaTypeEnum {
	return CreateEdgeWafDomainsRequestBodyAreaTypeEnum{
		MAINLAND_CHINA: CreateEdgeWafDomainsRequestBodyAreaType{
			value: "mainland_china",
		},
		OUTSIDE_MAINLAND_CHINA: CreateEdgeWafDomainsRequestBodyAreaType{
			value: "outside_mainland_china",
		},
		GLOBAL: CreateEdgeWafDomainsRequestBodyAreaType{
			value: "global",
		},
		EUROPE: CreateEdgeWafDomainsRequestBodyAreaType{
			value: "europe",
		},
	}
}

func (c CreateEdgeWafDomainsRequestBodyAreaType) Value() string {
	return c.value
}

func (c CreateEdgeWafDomainsRequestBodyAreaType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateEdgeWafDomainsRequestBodyAreaType) UnmarshalJSON(b []byte) error {
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
