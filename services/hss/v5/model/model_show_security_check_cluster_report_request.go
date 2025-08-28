package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSecurityCheckClusterReportRequest Request Object
type ShowSecurityCheckClusterReportRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// language
	XLanguage *ShowSecurityCheckClusterReportRequestXLanguage `json:"x-language,omitempty"`

	// **参数解释**： 集群的安全体检报告ID **约束限制**： 不涉及 **取值范围**： 字符长度1-64位 **默认取值**： 不涉及
	ReportId string `json:"report_id"`
}

func (o ShowSecurityCheckClusterReportRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityCheckClusterReportRequest struct{}"
	}

	return strings.Join([]string{"ShowSecurityCheckClusterReportRequest", string(data)}, " ")
}

type ShowSecurityCheckClusterReportRequestXLanguage struct {
	value string
}

type ShowSecurityCheckClusterReportRequestXLanguageEnum struct {
	ZH_CN ShowSecurityCheckClusterReportRequestXLanguage
	EN_US ShowSecurityCheckClusterReportRequestXLanguage
}

func GetShowSecurityCheckClusterReportRequestXLanguageEnum() ShowSecurityCheckClusterReportRequestXLanguageEnum {
	return ShowSecurityCheckClusterReportRequestXLanguageEnum{
		ZH_CN: ShowSecurityCheckClusterReportRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowSecurityCheckClusterReportRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowSecurityCheckClusterReportRequestXLanguage) Value() string {
	return c.value
}

func (c ShowSecurityCheckClusterReportRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSecurityCheckClusterReportRequestXLanguage) UnmarshalJSON(b []byte) error {
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
