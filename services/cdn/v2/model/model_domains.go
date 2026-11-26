package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Domains 域名信息
type Domains struct {

	// **参数解释：** 加速域名ID。 **约束限制：** 不涉及 **取值范围：** 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释：** 加速域名。 **约束限制：** 不涉及 **取值范围：** 不涉及
	DomainName *string `json:"domain_name,omitempty"`

	// 域名业务类型，若为web，则表示类型为网站加速；若为download，则表示业务类型为文件下载加速；若为video，则表示业务类型为点播加速；若为wholeSite，则表示类型为全站加速。
	BusinessType *string `json:"business_type,omitempty"`

	// 加速域名状态。取值意义： - online表示“已开启” - offline表示“已停用” - configuring表示“配置中” - configure_failed表示“配置失败” - checking表示“审核中” - check_failed表示“审核未通过” - deleting表示“删除中”。
	DomainStatus *string `json:"domain_status,omitempty"`

	// **参数解释：** 加速域名对应的CNAME。 **约束限制：** 不涉及 **取值范围：** 不涉及
	Cname *string `json:"cname,omitempty"`

	// 源站配置。
	Sources *[]Sources `json:"sources,omitempty"`

	DomainOriginHost *DomainOriginHost `json:"domain_origin_host,omitempty"`

	// 是否开启HTTPS加速。取值意义： - 0：未开启 - 1：开启
	HttpsStatus *int32 `json:"https_status,omitempty"`

	// **参数解释：** 域名创建时间，相对于UTC 1970-01-01到当前时间相隔的毫秒数。 **约束限制：** 不涉及 **取值范围：** 不涉及
	CreateTime *int64 `json:"create_time,omitempty"`

	// **参数解释：** 域名修改时间，相对于UTC 1970-01-01到当前时间相隔的毫秒数。 **约束限制：** 不涉及 **取值范围：** 不涉及
	ModifyTime *int64 `json:"modify_time,omitempty"`

	// 封禁状态（0代表未禁用；1代表禁用）。
	Disabled *int32 `json:"disabled,omitempty"`

	// 锁定状态（0代表未锁定；1代表锁定）。
	Locked *int32 `json:"locked,omitempty"`

	// 自动刷新预热（0代表关闭；1代表打开）。
	AutoRefreshPreheat *int32 `json:"auto_refresh_preheat,omitempty"`

	// 华为云CDN提供的加速服务范围，包含：mainland_china中国大陆、outside_mainland_china中国大陆境外、global全球。
	ServiceArea *DomainsServiceArea `json:"service_area,omitempty"`

	// Range回源状态。取值意义： - on：开启 - off：关闭
	RangeStatus *string `json:"range_status,omitempty"`

	// 回源跟随状态。取值意义： - on：开启 - off：关闭
	FollowStatus *string `json:"follow_status,omitempty"`

	// 是否暂停源站回源。取值意义： - on：开启 - off：关闭
	OriginStatus *string `json:"origin_status,omitempty"`

	// **参数解释：** 域名禁用原因。 **约束限制：** 不涉及 **取值范围：** 不涉及
	BannedReason *string `json:"banned_reason,omitempty"`

	// **参数解释：** 域名锁定原因。 **约束限制：** 不涉及 **取值范围：** 不涉及
	LockedReason *string `json:"locked_reason,omitempty"`

	// **参数解释：** 企业项目ID。 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id。 **约束限制：** 不涉及 **取值范围：** 不涉及
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 标签信息。
	Tags *[]EpResourceTag `json:"tags,omitempty"`

	// **参数解释：** 指定CDN回源时的请求协议 **约束限制：** 不涉及 **取值范围：** - follow: 协议跟随 - http: http协议 - https: https协议  **默认取值：** http: http协议
	OriginProtocol *string `json:"origin_protocol,omitempty"`
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
