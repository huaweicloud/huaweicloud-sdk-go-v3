package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RiskHandleInfoBaseLineInfo 基线信息
type RiskHandleInfoBaseLineInfo struct {

	// **参数解释**: 存在配置风险的主机数 **取值范围**: 最小值0，最大值2147483647
	ConfigRiskHostNum *int64 `json:"config_risk_host_num,omitempty"`

	// **参数解释**: 存在的配置风险数 **取值范围**: 最小值0，最大值2147483647
	ConfigRiskNum *int64 `json:"config_risk_num,omitempty"`

	// **参数解释**: 通过率 **取值范围**: 最小值0，最大值1
	PassedRate *float32 `json:"passed_rate,omitempty"`

	// **参数解释**: 通过率击败的用户率 **取值范围**: 最小值0，最大值1
	BeatRate *float32 `json:"beat_rate,omitempty"`

	// **参数解释**: 弱口令数 **取值范围**: 最小值0，最大值2147483647
	WeakPwdNum *int64 `json:"weak_pwd_num,omitempty"`

	// **参数解释**: 通过量 **取值范围**: 最小值0，最大值2147483647
	PassedNum *int64 `json:"passed_num,omitempty"`

	// **参数解释**: 总的风险数 **取值范围**: 最小值0，最大值2147483647
	TotalConfigRiskNum *int64 `json:"total_config_risk_num,omitempty"`

	// **参数解释**: 推荐版本，只支持企业版 hss.version.enterprise **取值范围**: 字符长度1-32位
	VersionRecommend *string `json:"version_recommend,omitempty"`
}

func (o RiskHandleInfoBaseLineInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RiskHandleInfoBaseLineInfo struct{}"
	}

	return strings.Join([]string{"RiskHandleInfoBaseLineInfo", string(data)}, " ")
}
