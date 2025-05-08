package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckScriptRiskResData 返回数据。
type CheckScriptRiskResData struct {

	// 风险等级。
	RiskLevel string `json:"risk_level"`

	// 黑名单列表。
	BlacklistCommands *[]CheckScriptRiskResDataBlacklistCommands `json:"blacklist_commands,omitempty"`
}

func (o CheckScriptRiskResData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckScriptRiskResData struct{}"
	}

	return strings.Join([]string{"CheckScriptRiskResData", string(data)}, " ")
}
