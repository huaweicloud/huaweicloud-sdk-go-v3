package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 域名信息
type Domains struct {

	// 加速域名ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 加速域名。
	DomainName *string `json:"domain_name,omitempty" xml:"domain_name"`

	// 域名业务类型，若为web，则表示类型为网站加速；若为download，则表示业务类型为文件下载加速；若为video，则表示业务类型为点播加速；若为wholeSite，则表示类型为全站加速。
	BusinessType *string `json:"business_type,omitempty" xml:"business_type"`

	// 域名所属用户的domain_id。
	UserDomainId *string `json:"user_domain_id,omitempty" xml:"user_domain_id"`

	// 加速域名状态。取值意义： - online表示“已开启” - offline表示“已停用” - configuring表示“配置中” - configure_failed表示“配置失败” - checking表示“审核中” - check_failed表示“审核未通过” - deleting表示“删除中”
	DomainStatus *string `json:"domain_status,omitempty" xml:"domain_status"`

	// 加速域名对应的CNAME。
	Cname *string `json:"cname,omitempty" xml:"cname"`

	// 源站域名或源站IP，源站为IP类型时，仅支持IPv4，如需传入多个源站IP，以多个源站对象传入，除IP其他参数请保持一致，主源站最多支持15个源站IP对象，备源站最多支持15个源站IP对象；源站为域名类型时仅支持1个源站对象。不支持IP源站和域名源站混用。
	Sources *[]Sources `json:"sources,omitempty" xml:"sources"`

	DomainOriginHost *DomainOriginHost `json:"domain_origin_host,omitempty" xml:"domain_origin_host"`

	// 是否开启HTTPS加速。
	HttpsStatus *int32 `json:"https_status,omitempty" xml:"https_status"`

	// 域名创建时间，相对于UTC 1970-01-01到当前时间相隔的毫秒数。
	CreateTime *int64 `json:"create_time,omitempty" xml:"create_time"`

	// 域名修改时间，相对于UTC 1970-01-01到当前时间相隔的毫秒数。
	ModifyTime *int64 `json:"modify_time,omitempty" xml:"modify_time"`

	// 封禁状态（0代表未禁用；1代表禁用）。
	Disabled *int32 `json:"disabled,omitempty" xml:"disabled"`

	// 锁定状态（0代表未锁定；1代表锁定）。
	Locked *int32 `json:"locked,omitempty" xml:"locked"`

	// 自动刷新预热（0代表关闭；1代表打开）
	AutoRefreshPreheat *int32 `json:"auto_refresh_preheat,omitempty" xml:"auto_refresh_preheat"`

	// 华为云CDN提供的加速服务范围，包含：mainland_china中国大陆、outside_mainland_china中国大陆境外、global全球。
	ServiceArea *DomainsServiceArea `json:"service_area,omitempty" xml:"service_area"`

	// Range回源状态。
	RangeStatus *string `json:"range_status,omitempty" xml:"range_status"`

	// 回源跟随状态。
	FollowStatus *string `json:"follow_status,omitempty" xml:"follow_status"`

	// 是否暂停源站回源。
	OriginStatus *string `json:"origin_status,omitempty" xml:"origin_status"`

	// 域名禁用原因
	BannedReason *string `json:"banned_reason,omitempty" xml:"banned_reason"`

	// 域名锁定原因
	LockedReason *string `json:"locked_reason,omitempty" xml:"locked_reason"`

	// 当用户开启企业项目功能时，该参数生效，表示查询资源所属项目，不传表示查询默认项目。注意：当使用子账号调用接口时，该参数必传。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`
}

func (o Domains) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Domains struct{}"
	}

	return strings.Join([]string{"Domains", string(data)}, " ")
}

type DomainsServiceArea struct {
	value string
}

type DomainsServiceAreaEnum struct {
	MAINLAND_CHINA         DomainsServiceArea
	OUTSIDE_MAINLAND_CHINA DomainsServiceArea
	GLOBAL                 DomainsServiceArea
}

func GetDomainsServiceAreaEnum() DomainsServiceAreaEnum {
	return DomainsServiceAreaEnum{
		MAINLAND_CHINA: DomainsServiceArea{
			value: "mainland_china",
		},
		OUTSIDE_MAINLAND_CHINA: DomainsServiceArea{
			value: "outside_mainland_china",
		},
		GLOBAL: DomainsServiceArea{
			value: "global",
		},
	}
}

func (c DomainsServiceArea) Value() string {
	return c.value
}

func (c DomainsServiceArea) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DomainsServiceArea) UnmarshalJSON(b []byte) error {
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
