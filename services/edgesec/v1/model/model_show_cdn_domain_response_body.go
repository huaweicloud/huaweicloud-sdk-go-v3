package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowCdnDomainResponseBody CDN域名详情
type ShowCdnDomainResponseBody struct {

	// 域名
	DomainName *string `json:"domain_name,omitempty"`

	// 加速域名状态。取值意义： - online表示“已开启” - offline表示“已停用” - configuring表示“配置中” - configure_failed表示“配置失败” - checking表示“审核中” - check_failed表示“审核未通过” - deleting表示“删除中”。
	DomainStatus *string `json:"domain_status,omitempty"`

	// 域名id
	DomainId *string `json:"domain_id,omitempty"`

	// 域名使用的证书id
	CertificateId *string `json:"certificate_id,omitempty"`

	// 域名业务区域
	ServiceArea *ShowCdnDomainResponseBodyServiceArea `json:"service_area,omitempty"`

	// 是否开启ipv6加速：0关闭/1开启
	Ipv6Accelerate *int32 `json:"ipv6_accelerate,omitempty"`

	// 域名业务类型。取值意义： - web表示“网站加速” - download表示“文件下载加速” - video表示“点播加速” - wholeSite表示“全站加速”
	BusinessType *string `json:"business_type,omitempty"`

	// 是否启用https：0关闭/1开启
	HttpsStatus *int32 `json:"https_status,omitempty"`

	// 强制重定向：0不开启重定向/1强制重定向为HTTP/2强制重定向为HTTPS
	ForceRedirect *int32 `json:"force_redirect,omitempty"`

	ExtendedTags *CdnDomainTags `json:"extended_tags,omitempty"`

	// 是否为waf防护域名
	IsAdded *bool `json:"is_added,omitempty"`
}

func (o ShowCdnDomainResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCdnDomainResponseBody struct{}"
	}

	return strings.Join([]string{"ShowCdnDomainResponseBody", string(data)}, " ")
}

type ShowCdnDomainResponseBodyServiceArea struct {
	value string
}

type ShowCdnDomainResponseBodyServiceAreaEnum struct {
	MAINLAND_CHINA         ShowCdnDomainResponseBodyServiceArea
	OUTSIDE_MAINLAND_CHINA ShowCdnDomainResponseBodyServiceArea
	GLOBAL                 ShowCdnDomainResponseBodyServiceArea
	EUROPE                 ShowCdnDomainResponseBodyServiceArea
}

func GetShowCdnDomainResponseBodyServiceAreaEnum() ShowCdnDomainResponseBodyServiceAreaEnum {
	return ShowCdnDomainResponseBodyServiceAreaEnum{
		MAINLAND_CHINA: ShowCdnDomainResponseBodyServiceArea{
			value: "mainland_china",
		},
		OUTSIDE_MAINLAND_CHINA: ShowCdnDomainResponseBodyServiceArea{
			value: "outside_mainland_china",
		},
		GLOBAL: ShowCdnDomainResponseBodyServiceArea{
			value: "global",
		},
		EUROPE: ShowCdnDomainResponseBodyServiceArea{
			value: "europe",
		},
	}
}

func (c ShowCdnDomainResponseBodyServiceArea) Value() string {
	return c.value
}

func (c ShowCdnDomainResponseBodyServiceArea) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowCdnDomainResponseBodyServiceArea) UnmarshalJSON(b []byte) error {
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
