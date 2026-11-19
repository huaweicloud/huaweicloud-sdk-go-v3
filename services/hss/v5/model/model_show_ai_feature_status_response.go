package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAiFeatureStatusResponse Response Object
type ShowAiFeatureStatusResponse struct {

	// **参数解释**: AI Defender白名单开关状态 **取值范围**: - not_enabled：该Region已上线AI Defender，但该客户未开通，需显示引导页面 - enabled：该Region已上线AI Defender，并且该客户已开通
	AiDefender *string `json:"ai_defender,omitempty"`

	// **参数解释**: AgentArts开关状态 **取值范围**: - not_available：该Region未上线AgentArts，不显示AgentArts任何内容 - not_enabled：该Region已上线AgentArts，但该客户未开通 - enabled：该Region已上线AgentArts，并且该客户已开通
	Agentarts *string `json:"agentarts,omitempty"`

	// **参数解释**: OfficeAce开关状态 **取值范围**: - not_available：该Region未上线OfficeAce，不显示OfficeAce任何内容 - not_enabled：该Region已上线OfficeAce，但该客户未开通 - enabled：该Region已上线OfficeAce，并且该客户已开通
	Officeace *string `json:"officeace,omitempty"`

	// **参数解释**: OpenClaw开关状态 **取值范围**: - not_available：该Region未上线OpenClaw - enabled：该Region已上线OpenClaw
	Openclaw       *string `json:"openclaw,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAiFeatureStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAiFeatureStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowAiFeatureStatusResponse", string(data)}, " ")
}
