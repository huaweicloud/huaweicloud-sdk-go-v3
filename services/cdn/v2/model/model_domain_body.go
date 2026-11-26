package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DomainBody **参数解释：** 创建域名请求体，包含创建域名所需参数 **约束限制：** 不涉及
type DomainBody struct {

	// **参数解释：** 加速域名 > 添加泛域名后，该泛域名所有次级域名默认支持CDN加速  **约束限制：** 加速域名不允许重复添加 **取值范围：** - 域名长度不能超过200个字符 - 支持大小写字母、数字、\"-\"、\".\"，首尾字符不能是\"-\"或\".\" - 泛域名场景下支持\"\\*\"，且\"\\*必须为首字符 - 域名单节点长度不超过63个字符，即：\\*\\*\\*.\\*\\*\\*.com中，\\*\\*\\*的字符数不超过63个字符  **默认取值：** 不涉及
	DomainName string `json:"domain_name"`

	// **参数解释：** 域名业务类型 **约束限制：** 不涉及 **取值范围：** - web: 网页加速 - download: 文件下载加速 - video: 点播加速 - wholeSite: 全站加速  **默认取值：** 不涉及
	BusinessType DomainBodyBusinessType `json:"business_type"`

	// **参数解释：** 源站配置 **约束限制：** 不涉及
	Sources []SourcesRequestBody `json:"sources"`

	// **参数解释：** 域名服务范围 **约束限制：** 服务范围为中国大陆或全球时，加速域名需要到工信部备案 **取值范围：** - mainland_china: 中国大陆 - outside_mainland_china: 中国大陆境外 - global: 全球  **默认取值：** mainland_china: 中国大陆
	ServiceArea *DomainBodyServiceArea `json:"service_area,omitempty"`

	// **参数解释：** 企业项目id。您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id **约束限制：** 当用户开启企业项目功能时，该参数生效，表示查询资源所属项目，\"all\"表示所有项目  **取值范围：** 不涉及 **默认取值：** 不涉及
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释：** 指定CDN回源时的请求协议 **约束限制：** 不涉及 **取值范围：** - follow: 协议跟随 - http: http协议 - https: https协议  **默认取值：** http: http协议
	OriginProtocol *string `json:"origin_protocol,omitempty"`

	Configs *CreateDomainConfigs `json:"configs,omitempty"`
}

func (o DomainBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainBody struct{}"
	}

	return strings.Join([]string{"DomainBody", string(data)}, " ")
}

type DomainBodyBusinessType struct {
	value string
}

type DomainBodyBusinessTypeEnum struct {
	WEB        DomainBodyBusinessType
	DOWNLOAD   DomainBodyBusinessType
	VIDEO      DomainBodyBusinessType
	WHOLE_SITE DomainBodyBusinessType
}

func GetDomainBodyBusinessTypeEnum() DomainBodyBusinessTypeEnum {
	return DomainBodyBusinessTypeEnum{
		WEB: DomainBodyBusinessType{
			value: "web",
		},
		DOWNLOAD: DomainBodyBusinessType{
			value: "download",
		},
		VIDEO: DomainBodyBusinessType{
			value: "video",
		},
		WHOLE_SITE: DomainBodyBusinessType{
			value: "wholeSite",
		},
	}
}

func (c DomainBodyBusinessType) Value() string {
	return c.value
}

func (c DomainBodyBusinessType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DomainBodyBusinessType) UnmarshalJSON(b []byte) error {
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

type DomainBodyServiceArea struct {
	value string
}

type DomainBodyServiceAreaEnum struct {
	MAINLAND_CHINA         DomainBodyServiceArea
	OUTSIDE_MAINLAND_CHINA DomainBodyServiceArea
	GLOBAL                 DomainBodyServiceArea
}

func GetDomainBodyServiceAreaEnum() DomainBodyServiceAreaEnum {
	return DomainBodyServiceAreaEnum{
		MAINLAND_CHINA: DomainBodyServiceArea{
			value: "mainland_china",
		},
		OUTSIDE_MAINLAND_CHINA: DomainBodyServiceArea{
			value: "outside_mainland_china",
		},
		GLOBAL: DomainBodyServiceArea{
			value: "global",
		},
	}
}

func (c DomainBodyServiceArea) Value() string {
	return c.value
}

func (c DomainBodyServiceArea) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DomainBodyServiceArea) UnmarshalJSON(b []byte) error {
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
