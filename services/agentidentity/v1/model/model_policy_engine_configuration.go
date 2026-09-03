package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyEngineConfiguration TokenVault 关联的策略集配置。
type PolicyEngineConfiguration struct {

	// System-generated unique identifier for the policy engine.
	PolicyEngineId *string `json:"policy_engine_id,omitempty"`

	Mode *PolicyEngineMode `json:"mode,omitempty"`
}

func (o PolicyEngineConfiguration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyEngineConfiguration struct{}"
	}

	return strings.Join([]string{"PolicyEngineConfiguration", string(data)}, " ")
}
