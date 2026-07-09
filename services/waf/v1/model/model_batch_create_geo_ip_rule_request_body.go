package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchCreateGeoIpRuleRequestBody struct {

	// 地理位置控制规则名称
	Name *string `json:"name,omitempty"`

	// 地理位置封禁区域，可通过ShowPolicyGeoipMap接口查询支持的区域
	Geoip string `json:"geoip"`

	// 防护动作：  - 0 拦截  - 1 放行  - 2 仅记录
	White int32 `json:"white"`

	// ip范围。若您的网站使用独享模式，请确认独享引擎是否全部升级到最新版本，避免造成异常。202412之后的版本支持配置IP范围
	IpType *string `json:"ip_type,omitempty"`

	// 规则描述
	Description *string `json:"description,omitempty"`

	// 需要添加规则的策略id
	PolicyIds []string `json:"policy_ids"`
}

func (o BatchCreateGeoIpRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateGeoIpRuleRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateGeoIpRuleRequestBody", string(data)}, " ")
}
