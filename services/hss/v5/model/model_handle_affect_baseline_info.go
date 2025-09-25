package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HandleAffectBaselineInfo 该操作影响范围的详细信息
type HandleAffectBaselineInfo struct {

	// **参数解释** 主机id **取值范围** 字符长度1-256位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释** 服务器名称 **取值范围** 字符长度1-64位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释** 服务器公网ip **取值范围** 字符长度0-128位
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释** 服务器私网ip **取值范围** 字符长度0-2048位
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释** 资产重要性，包含如下3种 **取值范围** - important ：重要资产 - common    ：一般资产 - test      ：测试资产
	AssetValue *string `json:"asset_value,omitempty"`

	// **参数解释** 基线检查的基线名称 **取值范围** 字符长度0-255位
	CheckType *string `json:"check_type,omitempty"`

	// **参数解释** 标准类型，包含如下3种 **取值范围** - cn_standard : 等保合规标准 - hw_standard : 云安全实践标准 - cis_standard : 通用安全标准
	Standard *string `json:"standard,omitempty"`

	// **参数解释** 基线检查中检查项的检查类型 **取值范围** 字符长度0-128位
	Tag *string `json:"tag,omitempty"`

	// **参数解释** 基线检查中检查项的名称 **取值范围** 字符长度0-2048位
	CheckRuleName *string `json:"check_rule_name,omitempty"`
}

func (o HandleAffectBaselineInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HandleAffectBaselineInfo struct{}"
	}

	return strings.Join([]string{"HandleAffectBaselineInfo", string(data)}, " ")
}
