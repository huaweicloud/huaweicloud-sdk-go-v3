package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 域名信息
type DomainsWithPort struct {

	// 加速域名ID。
	Id *string `json:"id,omitempty"`

	// 域名业务类型，若为web，则表示类型为网站加速；若为download，则表示业务类型为文件下载加速；若为video，则表示业务类型为点播加速；若为wholeSite，则表示类型为全站加速。
	BusinessType *string `json:"business_type,omitempty"`

	// 加速域名状态。取值意义： - online表示“已开启” - offline表示“已停用” - configuring表示“配置中” - configure_failed表示“配置失败” - checking表示“审核中” - check_failed表示“审核未通过” - deleting表示“删除中”。
	DomainStatus *string `json:"domain_status,omitempty"`

	// 加速域名对应的CNAME。
	Cname *string `json:"cname,omitempty"`

	// 源站域名或源站IP，源站为IP类型时，仅支持IPv4，如需传入多个源站IP，以多个源站对象传入，除IP其他参数请保持一致，主源站最多支持15个源站IP对象，备源站最多支持15个源站IP对象；源站为域名类型时仅支持1个源站对象。不支持IP源站和域名源站混用。
	Sources *[]SourcesConfig `json:"sources,omitempty"`

	// 是否开启HTTPS加速。 0：代表未开启HTTPS加速； 1：代表开启HTTPS加速，且回源方式为协议跟随； 2：代表开启HTTPS加速，且回源方式为HTTP； 3：代表开启HTTPS加速，且回源方式为HTTPS。
	HttpsStatus *int32 `json:"https_status,omitempty"`

	// 域名创建时间，相对于UTC 1970-01-01到当前时间相隔的毫秒数。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 域名修改时间，相对于UTC 1970-01-01到当前时间相隔的毫秒数。
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 封禁状态（0代表未禁用；1代表禁用）。
	Disabled *int32 `json:"disabled,omitempty"`

	// 锁定状态（0代表未锁定；1代表锁定）。
	Locked *int32 `json:"locked,omitempty"`

	// 华为云CDN提供的加速服务范围，包含：mainland_china中国大陆、outside_mainland_china中国大陆境外、global全球。
	ServiceArea *DomainsWithPortServiceArea `json:"service_area,omitempty"`
}

func (o DomainsWithPort) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainsWithPort struct{}"
	}

	return strings.Join([]string{"DomainsWithPort", string(data)}, " ")
}

type DomainsWithPortServiceArea struct {
	value string
}

type DomainsWithPortServiceAreaEnum struct {
	MAINLAND_CHINA         DomainsWithPortServiceArea
	OUTSIDE_MAINLAND_CHINA DomainsWithPortServiceArea
	GLOBAL                 DomainsWithPortServiceArea
}

func GetDomainsWithPortServiceAreaEnum() DomainsWithPortServiceAreaEnum {
	return DomainsWithPortServiceAreaEnum{
		MAINLAND_CHINA: DomainsWithPortServiceArea{
			value: "mainland_china",
		},
		OUTSIDE_MAINLAND_CHINA: DomainsWithPortServiceArea{
			value: "outside_mainland_china",
		},
		GLOBAL: DomainsWithPortServiceArea{
			value: "global",
		},
	}
}

func (c DomainsWithPortServiceArea) Value() string {
	return c.value
}

func (c DomainsWithPortServiceArea) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DomainsWithPortServiceArea) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
