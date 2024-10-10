package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FrequencyControlRule struct {

	// id
	Id *string `json:"id,omitempty"`

	// 判断是否是智能cc生成的规则
	Producer *int32 `json:"producer,omitempty"`

	// 规则名称
	Name *string `json:"name,omitempty"`

	// 规则应用的url
	Url *string `json:"url,omitempty"`

	// 限速频率，单位为次，范围为1~2147483647
	LimitNum *string `json:"limit_num,omitempty"`

	// 限速周期，单位为秒，范围1~3600
	LimitPeriod *string `json:"limit_period,omitempty"`

	// 阻断时间，单位为秒，范围为0~65535
	LockTime *string `json:"lock_time,omitempty"`

	// 限速模式：ip、cookie、header、other、policy、domain、url。 源限速：ip：IP限速，根据IP区分单个Web访问者。cookie：用户限速，根据Cookie键值区分单个Web访问者。header：用户限速，根据Header区分单个Web访问者。other：根据Referer（自定义请求访问的来源）字段区分单个Web访问者。-目的限速：policy: 策略限速、domain: 域名限速、url: url限速
	TagType *string `json:"tag_type,omitempty"`

	// 用户标识，当限速模式为用户限速(cookie或header)时
	TagIndex *string `json:"tag_index,omitempty"`

	TagCondition *TagCondition `json:"tag_condition,omitempty"`

	Action *ActionInfo `json:"action,omitempty"`

	// cc规则防护模式，0：标准(老版本)，只支持对域名的防护路径做限制。1：高级(新版本)，支持对路径、IP、Cookie、Header、Params字段做限制。修改CC规则时必须传mode
	Mode *string `json:"mode,omitempty"`

	// cc规则防护规则限速条件
	Conditions *[]Condition `json:"conditions,omitempty"`

	// 放行频率，单位为次，范围为0~2147483647
	UnlockNum *int32 `json:"unlock_num,omitempty"`

	// 域名聚合统计
	DomainAggregation *bool `json:"domain_aggregation,omitempty"`

	// 全局计数
	RegionAggregation *bool `json:"region_aggregation,omitempty"`

	// 锁定验证时间
	CaptchaLockTime *int32 `json:"captcha_lock_time,omitempty"`

	// 是否灰度生效
	GrayscaleTime *bool `json:"grayscale_time,omitempty"`
}

func (o FrequencyControlRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FrequencyControlRule struct{}"
	}

	return strings.Join([]string{"FrequencyControlRule", string(data)}, " ")
}
