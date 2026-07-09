package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGeoipRuleResponse Response Object
type DeleteGeoipRuleResponse struct {

	// 规则id
	Id *string `json:"id,omitempty"`

	// 地理位置控制规则名称
	Name *string `json:"name,omitempty"`

	// 策略id
	Policyid *string `json:"policyid,omitempty"`

	// 地理位置封禁区域
	Geoip *string `json:"geoip,omitempty"`

	// 地理位置封禁区域
	GeoTagList *[]string `json:"geoTagList,omitempty"`

	// 防护动作：  - 0 拦截  - 1 放行  - 2 仅记录
	White *int32 `json:"white,omitempty"`

	// **参数解释：** 规则状态标识，用于指定规则的启用或关闭状态 **约束限制：** 不涉及 **取值范围：**  - 0：关闭  - 1：开启 **默认取值：** 不涉及
	Status *int32 `json:"status,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 创建规则时间戳
	Timestamp      *int64 `json:"timestamp,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o DeleteGeoipRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGeoipRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteGeoipRuleResponse", string(data)}, " ")
}
