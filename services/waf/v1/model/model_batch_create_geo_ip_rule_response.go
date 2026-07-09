package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateGeoIpRuleResponse Response Object
type BatchCreateGeoIpRuleResponse struct {

	// 规则id
	Id *string `json:"id,omitempty"`

	// 地理位置控制规则名称
	Name *string `json:"name,omitempty"`

	// 策略id
	Policyid *string `json:"policyid,omitempty"`

	// 地理位置封禁区域
	Geoip *string `json:"geoip,omitempty"`

	// 防护动作：  - 0 拦截  - 1 放行  - 2 仅记录
	White *int32 `json:"white,omitempty"`

	// **参数解释：** 规则状态标识，用于指定规则的启用或关闭状态 **约束限制：** 不涉及 **取值范围：**  - 0：关闭  - 1：开启 **默认取值：** 不涉及
	Status *int32 `json:"status,omitempty"`

	// ip范围。若您的网站使用独享模式，请确认独享引擎是否全部升级到最新版本，避免造成异常。202412之后的版本支持配置IP范围
	IpType *string `json:"ip_type,omitempty"`

	// 规则描述
	Description *string `json:"description,omitempty"`

	// 创建规则时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 地理位置
	GeoTagList     *[]string `json:"geoTagList,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchCreateGeoIpRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateGeoIpRuleResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateGeoIpRuleResponse", string(data)}, " ")
}
