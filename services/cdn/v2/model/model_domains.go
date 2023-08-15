package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Domains 域名信息
type Domains struct {

	// 加速域名ID。
	Id *string `json:"id,omitempty"`

	// 加速域名。
	DomainName *string `json:"domain_name,omitempty"`

	// 域名业务类型，若为web，则表示类型为网站加速；若为download，则表示业务类型为文件下载加速；若为video，则表示业务类型为点播加速；若为wholeSite，则表示类型为全站加速。
	BusinessType *string `json:"business_type,omitempty"`

	// 加速域名状态。取值意义： - online表示“已开启” - offline表示“已停用” - configuring表示“配置中” - configure_failed表示“配置失败” - checking表示“审核中” - check_failed表示“审核未通过” - deleting表示“删除中”。
	DomainStatus *string `json:"domain_status,omitempty"`

	// 加速域名对应的CNAME。
	Cname *string `json:"cname,omitempty"`

	// 源站配置。
	Sources *[]Sources `json:"sources,omitempty"`

	DomainOriginHost *DomainOriginHost `json:"domain_origin_host,omitempty"`

	// 是否开启HTTPS加速。
	HttpsStatus *int32 `json:"https_status,omitempty"`

	// 域名创建时间，相对于UTC 1970-01-01到当前时间相隔的毫秒数。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 域名修改时间，相对于UTC 1970-01-01到当前时间相隔的毫秒数。
	ModifyTime *int64 `json:"modify_time,omitempty"`

	// 封禁状态（0代表未禁用；1代表禁用）。
	Disabled *int32 `json:"disabled,omitempty"`

	// 锁定状态（0代表未锁定；1代表锁定）。
	Locked *int32 `json:"locked,omitempty"`

	// 自动刷新预热（0代表关闭；1代表打开）。
	AutoRefreshPreheat *int32 `json:"auto_refresh_preheat,omitempty"`

	// 华为云CDN提供的加速服务范围，包含：mainland_china中国大陆、outside_mainland_china中国大陆境外、global全球。
	ServiceArea *DomainsServiceArea `json:"service_area,omitempty"`

	// Range回源状态。
	RangeStatus *string `json:"range_status,omitempty"`

	// 回源跟随状态。
	FollowStatus *string `json:"follow_status,omitempty"`

	// 是否暂停源站回源（off代表关闭 on代表开启）。
	OriginStatus *string `json:"origin_status,omitempty"`

	// 域名禁用原因。
	BannedReason *string `json:"banned_reason,omitempty"`

	// 域名锁定原因。
	LockedReason *string `json:"locked_reason,omitempty"`

	// 当用户开启企业项目功能时，该参数生效，表示查询资源所属项目，不传表示查询默认项目。注意：当使用子帐号调用接口时，该参数必传。  您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 标签信息。
	Tags *[]EpResourceTag `json:"tags,omitempty"`
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
