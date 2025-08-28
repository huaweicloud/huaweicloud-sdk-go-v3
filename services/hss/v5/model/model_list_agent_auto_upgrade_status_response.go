package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAgentAutoUpgradeStatusResponse Response Object
type ListAgentAutoUpgradeStatusResponse struct {

	// **参数解释**： Agent自动升级配置开关状态 **取值范围**： - true：开启 - false：关闭
	Enabled        *bool `json:"enabled,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ListAgentAutoUpgradeStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAgentAutoUpgradeStatusResponse struct{}"
	}

	return strings.Join([]string{"ListAgentAutoUpgradeStatusResponse", string(data)}, " ")
}
