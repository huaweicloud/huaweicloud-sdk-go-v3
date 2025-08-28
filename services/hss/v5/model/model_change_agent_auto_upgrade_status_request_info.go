package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeAgentAutoUpgradeStatusRequestInfo Agent自动升级配置开关状态请求体
type ChangeAgentAutoUpgradeStatusRequestInfo struct {

	// **参数解释**： Agent自动升级配置开关状态 **约束限制**: 不涉及 **取值范围**： - true：开启 - false：关闭  **默认取值**: false
	Enabled *bool `json:"enabled,omitempty"`
}

func (o ChangeAgentAutoUpgradeStatusRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeAgentAutoUpgradeStatusRequestInfo struct{}"
	}

	return strings.Join([]string{"ChangeAgentAutoUpgradeStatusRequestInfo", string(data)}, " ")
}
