package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListWebBasicProtectionRulesRequest Request Object
type ListWebBasicProtectionRulesRequest struct {

	// **参数解释：** 语言，默认值为zh-cn。zh-cn（中文）/en-us（英文）。 **约束限制：** 不涉及 **取值范围：** - zh-cn：中文 - en-us：英文  **默认取值：** - zh-cn
	XLanguage *ListWebBasicProtectionRulesRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释：** 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目ID。若需要查询当前用户所有企业项目绑定的资源信息，请传参all_granted_eps。 **约束限制：** 不涉及 **取值范围：**  - 0：代表default企业项目  - all_granted_eps：代表所有企业项目  - 其它企业项目ID：长度为36个字符  **默认取值：** 0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释：** 分页查询的起始位置，表示从第几条记录开始返回（从1开始计数）。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 1
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 分页查询的单页返回数量，控制每次请求返回的记录条数。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 起始时间（13位毫秒时间戳），需要和to同时使用。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	From *int64 `json:"from,omitempty"`

	// **参数解释：** 结束时间（13位毫秒时间戳），需要和from同时使用。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	To *int64 `json:"to,omitempty"`

	// **参数解释：** 规则集的防护严格程度。规则集（宽松）下对业务的误报率降低，但漏报率可能会增高；而规则集（严格）下对业务的误报率可能会增高，但漏报率降低。 **约束限制：** 不涉及 **取值范围：** - 1：宽松 - 2：中等 - 3：严格  **默认取值：** 不涉及
	Level *ListWebBasicProtectionRulesRequestLevel `json:"level,omitempty"`

	// **参数解释：** 规则ID，规则的唯一标识。 **约束限制：** 不涉及 **取值范围：** 长度为6个字符 **默认取值：** 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释：** 规则描述 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释：** CVE编号 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	CveNumber *string `json:"cve_number,omitempty"`

	// **参数解释：** 危险等级 **约束限制：** 不涉及 **取值范围：** - 1：高危 - 2：中危 - 3：低危  **默认取值：** 不涉及
	RiskLevel *ListWebBasicProtectionRulesRequestRiskLevel `json:"risk_level,omitempty"`

	// **参数解释：** 防护类型 **约束限制：** 不涉及 **取值范围：** - vuln：其他 - xss：跨站脚本 - cmdi：命令注入 - lfi：本地文件包含 - rfi：远程文件包含 - webshell：网站木马 - robot：恶意爬虫 - sqli：SQL注入  **默认取值：** 不涉及
	ProtectionTypeNames *ListWebBasicProtectionRulesRequestProtectionTypeNames `json:"protection_type_names,omitempty"`

	// **参数解释：** 应用类型 **约束限制：** 不涉及 **取值范围：** 请参见WAF控制台，Web基础防护规则详情页面的应用类型。 **默认取值：** 不涉及
	ApplicationTypeNames *string `json:"application_type_names,omitempty"`
}

func (o ListWebBasicProtectionRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWebBasicProtectionRulesRequest struct{}"
	}

	return strings.Join([]string{"ListWebBasicProtectionRulesRequest", string(data)}, " ")
}

type ListWebBasicProtectionRulesRequestXLanguage struct {
	value string
}

type ListWebBasicProtectionRulesRequestXLanguageEnum struct {
	ZH_CN ListWebBasicProtectionRulesRequestXLanguage
	EN_US ListWebBasicProtectionRulesRequestXLanguage
}

func GetListWebBasicProtectionRulesRequestXLanguageEnum() ListWebBasicProtectionRulesRequestXLanguageEnum {
	return ListWebBasicProtectionRulesRequestXLanguageEnum{
		ZH_CN: ListWebBasicProtectionRulesRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListWebBasicProtectionRulesRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListWebBasicProtectionRulesRequestXLanguage) Value() string {
	return c.value
}

func (c ListWebBasicProtectionRulesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListWebBasicProtectionRulesRequestXLanguage) UnmarshalJSON(b []byte) error {
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

type ListWebBasicProtectionRulesRequestLevel struct {
	value int32
}

type ListWebBasicProtectionRulesRequestLevelEnum struct {
	E_1 ListWebBasicProtectionRulesRequestLevel
	E_2 ListWebBasicProtectionRulesRequestLevel
	E_3 ListWebBasicProtectionRulesRequestLevel
}

func GetListWebBasicProtectionRulesRequestLevelEnum() ListWebBasicProtectionRulesRequestLevelEnum {
	return ListWebBasicProtectionRulesRequestLevelEnum{
		E_1: ListWebBasicProtectionRulesRequestLevel{
			value: 1,
		}, E_2: ListWebBasicProtectionRulesRequestLevel{
			value: 2,
		}, E_3: ListWebBasicProtectionRulesRequestLevel{
			value: 3,
		},
	}
}

func (c ListWebBasicProtectionRulesRequestLevel) Value() int32 {
	return c.value
}

func (c ListWebBasicProtectionRulesRequestLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListWebBasicProtectionRulesRequestLevel) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type ListWebBasicProtectionRulesRequestRiskLevel struct {
	value int32
}

type ListWebBasicProtectionRulesRequestRiskLevelEnum struct {
	E_1 ListWebBasicProtectionRulesRequestRiskLevel
	E_2 ListWebBasicProtectionRulesRequestRiskLevel
	E_3 ListWebBasicProtectionRulesRequestRiskLevel
}

func GetListWebBasicProtectionRulesRequestRiskLevelEnum() ListWebBasicProtectionRulesRequestRiskLevelEnum {
	return ListWebBasicProtectionRulesRequestRiskLevelEnum{
		E_1: ListWebBasicProtectionRulesRequestRiskLevel{
			value: 1,
		}, E_2: ListWebBasicProtectionRulesRequestRiskLevel{
			value: 2,
		}, E_3: ListWebBasicProtectionRulesRequestRiskLevel{
			value: 3,
		},
	}
}

func (c ListWebBasicProtectionRulesRequestRiskLevel) Value() int32 {
	return c.value
}

func (c ListWebBasicProtectionRulesRequestRiskLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListWebBasicProtectionRulesRequestRiskLevel) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type ListWebBasicProtectionRulesRequestProtectionTypeNames struct {
	value string
}

type ListWebBasicProtectionRulesRequestProtectionTypeNamesEnum struct {
	VULN     ListWebBasicProtectionRulesRequestProtectionTypeNames
	XSS      ListWebBasicProtectionRulesRequestProtectionTypeNames
	CMDI     ListWebBasicProtectionRulesRequestProtectionTypeNames
	LFI      ListWebBasicProtectionRulesRequestProtectionTypeNames
	RFI      ListWebBasicProtectionRulesRequestProtectionTypeNames
	WEBSHELL ListWebBasicProtectionRulesRequestProtectionTypeNames
	ROBOT    ListWebBasicProtectionRulesRequestProtectionTypeNames
	SQLI     ListWebBasicProtectionRulesRequestProtectionTypeNames
}

func GetListWebBasicProtectionRulesRequestProtectionTypeNamesEnum() ListWebBasicProtectionRulesRequestProtectionTypeNamesEnum {
	return ListWebBasicProtectionRulesRequestProtectionTypeNamesEnum{
		VULN: ListWebBasicProtectionRulesRequestProtectionTypeNames{
			value: "vuln",
		},
		XSS: ListWebBasicProtectionRulesRequestProtectionTypeNames{
			value: "xss",
		},
		CMDI: ListWebBasicProtectionRulesRequestProtectionTypeNames{
			value: "cmdi",
		},
		LFI: ListWebBasicProtectionRulesRequestProtectionTypeNames{
			value: "lfi",
		},
		RFI: ListWebBasicProtectionRulesRequestProtectionTypeNames{
			value: "rfi",
		},
		WEBSHELL: ListWebBasicProtectionRulesRequestProtectionTypeNames{
			value: "webshell",
		},
		ROBOT: ListWebBasicProtectionRulesRequestProtectionTypeNames{
			value: "robot",
		},
		SQLI: ListWebBasicProtectionRulesRequestProtectionTypeNames{
			value: "sqli",
		},
	}
}

func (c ListWebBasicProtectionRulesRequestProtectionTypeNames) Value() string {
	return c.value
}

func (c ListWebBasicProtectionRulesRequestProtectionTypeNames) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListWebBasicProtectionRulesRequestProtectionTypeNames) UnmarshalJSON(b []byte) error {
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
